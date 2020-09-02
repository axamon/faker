// Package faker is a random data generator and struct fake data generator.
package faker

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	tagName   = "faker"
	skipTag   = "-"
	uniqueTag = "unique"
)

var (
	tagFnCallRegexp = regexp.MustCompile(`(.+?)\((.+?)\)`)
	tagLenRegexp    = regexp.MustCompile(`len=(\d+)`)
)

type fakerTag struct {
	funcName    string
	uniqueGroup string
	length      int
	params      []string
}

func (tag *fakerTag) mustSkip() bool {
	return tag.funcName == skipTag
}

// func decodeTag(tagString string) *fakerTag {
func decodeTag(structReflectType reflect.Type, fieldIndex int) *fakerTag {
	fieldReflectType := structReflectType.Field(fieldIndex)
	tagString := fieldReflectType.Tag.Get(tagName)
	tag := &fakerTag{}
	for _, token := range strings.Split(tagString, ";") {
		if token == skipTag {
			tag.funcName = skipTag
			return tag
		}
		if token == uniqueTag {
			tag.uniqueGroup = fmt.Sprintf("%s-%s", structReflectType.Name(), fieldReflectType.Name)
			continue
		}
		if m := tagLenRegexp.FindStringSubmatch(token); len(m) == 2 {
			tag.length, _ = strconv.Atoi(m[1])
			continue
		}
		if tag.funcName == "" {
			if m := tagFnCallRegexp.FindStringSubmatch(token); len(m) == 3 {
				tag.funcName = m[1]
				tag.params = strings.Split(m[2], ",")
				continue
			}
			tag.funcName = token
		}
	}
	return tag
}

// Build fills in exported elements of a struct with random data based on the
// value of `faker` tag of exported elements. The faker tag value can be any
// available function (case insensitive). Use `faker:"-"` to explicitly skip
// an element. Use `faker:"unique"` to guarantee a unique value. Use
// `faker:"len=x"` to specify the length of a slice or the size of a map (if
// ommitted a slice or a map with random size between 1 and 8 will be
// generated). Built-in types supported are: bool, int, int8, int16, int32,
// int64, uint, uint8, uint16, uint32, uint64, float32, float64, string. Other
// standard library supported types are time.Time and time.Duration. But is
// really easy to extend faker to add other builders to support other types
// and or customize faker's behavior (see RegisterBuilder function).
func Build(input interface{}) error {
	inputReflectType := reflect.TypeOf(input)
	if inputReflectType == nil {
		return errors.New("faker.Build input interface{} not allowed")
	}
	if inputReflectType.Kind() != reflect.Ptr {
		return errors.New("faker.Build input is not a pointer")
	}

	var err error
	inputReflectValue := reflect.ValueOf(input)
	if inputReflectType.Elem().Kind() == reflect.Slice {
		err = buildSlice(inputReflectValue.Elem(), &fakerTag{})
	} else {
		err = build(inputReflectValue, &fakerTag{})
	}
	if err != nil {
		return err
	}
	return nil
}

func buildSlice(inputReflectValue reflect.Value, tag *fakerTag) error {
	for i := 0; i < inputReflectValue.Len(); i++ {
		err := build(inputReflectValue.Index(i), tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func build(inputReflectValue reflect.Value, tag *fakerTag) error {
	inputReflectType := inputReflectValue.Type()
	kind := inputReflectType.Kind()

	var (
		fn    builderFunc
		found bool
		key   string
	)

	key = builderKey(tag.funcName, inputReflectType.String())
	fn, found = builders[key]

	if found {
		if !inputReflectValue.IsZero() {
			return nil
		}
		var (
			value interface{}
			err   error
		)
		if tag.uniqueGroup != "" {
			value, err = Uniq(tag.uniqueGroup, 0, func() (interface{}, error) {
				return fn(tag.params...)
			})
		} else {
			value, err = fn(tag.params...)
		}
		if err != nil {
			return err
		}
		inputReflectValue.Set(reflect.ValueOf(value))
		return nil
	}

	switch kind {
	case reflect.Ptr:
		if inputReflectValue.IsNil() {
			newVar := reflect.New(inputReflectType.Elem())
			inputReflectValue.Set(newVar)
		}
		return build(inputReflectValue.Elem(), tag)
	case reflect.Struct:
		for i := 0; i < inputReflectValue.NumField(); i++ {
			fieldTag := decodeTag(inputReflectType, i)
			if fieldTag.mustSkip() {
				continue
			}
			if !inputReflectValue.Field(i).CanSet() {
				continue // to avoid panic to set on unexported field in struct
			}
			if inputReflectValue.Field(i).Kind() == reflect.Ptr && inputReflectValue.Field(i).Type().Elem() == inputReflectType {
				continue // do not enter in an infinite loop on recursive type
			}
			err := build(inputReflectValue.Field(i), fieldTag)
			if err != nil {
				return err
			}
		}
		return nil
	case reflect.Slice:
		if inputReflectValue.IsNil() {
			var sliceLen int
			if tag != nil && tag.length != 0 {
				sliceLen = tag.length
			} else {
				sliceLen = IntInRange(1, 8)
			}
			newSlice := reflect.MakeSlice(inputReflectType, sliceLen, sliceLen)
			inputReflectValue.Set(newSlice)
			return buildSlice(inputReflectValue, tag)
		}
	case reflect.Map:
		if inputReflectValue.IsNil() {
			var (
				mapLen int
				key    reflect.Value
				elem   reflect.Value
				err    error
			)
			if tag != nil && tag.length != 0 {
				mapLen = tag.length
			} else {
				mapLen = IntInRange(1, 8)
			}
			keyReflectType := inputReflectType.Key()
			elemReflectType := inputReflectType.Elem()
			newMap := reflect.MakeMap(inputReflectType)
			for i := 0; i < mapLen; i++ {
				key = reflect.New(keyReflectType).Elem()
				elem = reflect.New(elemReflectType).Elem()
				err = build(key, tag)
				if err != nil {
					return err
				}
				err = build(elem, tag)
				if err != nil {
					return err
				}
				newMap.SetMapIndex(key, elem)
			}
			inputReflectValue.Set(newMap)
		}
	default:
		if tag.funcName != "" {
			return fmt.Errorf("invalid faker function '%s' for type '%s'", tag.funcName, inputReflectType.String())
		}
		return nil
	}
	return nil
}
