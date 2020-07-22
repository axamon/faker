package factory

import (
	"math"
)

func IntInRange(min, max int) int {
	if min >= max {
		return min
	}
	return random.Intn(max-min+1) + min
}

func Int() int {
	return IntInRange(math.MinInt32, math.MaxInt32)
}

func Int64InRange(min, max int64) int64 {
	if min >= max {
		return min
	}

	return random.Int63n(max-min) + min
}

func Int64() int64 {
	return random.Int63n(math.MaxInt64) + math.MinInt64
}

func Int32InRange(min, max int32) int32 {
	return int32(Int64InRange(int64(min), int64(max)))
}

func Int32() int32 {
	return Int32InRange(math.MinInt32, math.MaxInt32)
}

func Int16InRange(min, max int16) int16 {
	return int16(Int64InRange(int64(min), int64(max)))
}

func Int16() int16 {
	return Int16InRange(math.MinInt16, math.MaxInt16)
}

func Int8InRange(min, max int8) int8 {
	return int8(Int64InRange(int64(min), int64(max)))
}

func Int8() int8 {
	return Int8InRange(math.MinInt8, math.MaxInt8)
}

func UintInRange(min, max uint) uint {
	if min >= max {
		return min
	}
	return uint(random.Intn(int(max)-int(min)+1) + int(min))
}

func Uint() uint {
	return uint(IntInRange(0, math.MaxUint32))
}

func Uint64InRange(min, max uint64) uint64 {
	if min >= max {
		return min
	}
	return uint64(random.Int63n(int64(max)-int64(min)) + int64(min))
}

func Uint64() uint64 {
	return Uint64InRange(0, math.MaxInt64) + Uint64InRange(0, math.MaxInt64)
}

func Uint32InRange(min, max uint32) uint32 {
	return uint32(Uint64InRange(uint64(min), uint64(max)))
}

func Uint32() uint32 {
	return Uint32InRange(0, math.MaxUint32)
}

func Uint16InRange(min, max uint16) uint16 {
	return uint16(Uint64InRange(uint64(min), uint64(max)))
}

func Uint16() uint16 {
	return Uint16InRange(0, math.MaxUint16)
}

func Uint8InRange(min, max uint8) uint8 {
	return uint8(Uint64InRange(uint64(min), uint64(max)))
}

func Uint8() uint8 {
	return Uint8InRange(0, math.MaxUint8)
}

func Float64InRange(min, max float64) float64 {
	if min >= max {
		return min
	}
	return random.Float64()*(max-min) + min
}

func Float64() float64 {
	return Float64InRange(math.SmallestNonzeroFloat64, math.MaxFloat64)
}

func Float32InRange(min, max float32) float32 {
	if min >= max {
		return min
	}
	return random.Float32()*(max-min) + min
}

func Float32() float32 {
	return Float32InRange(math.SmallestNonzeroFloat32, math.MaxFloat32)
}

// Provider functions

func intInRangeFn(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return IntInRange(min, max), nil
}

func intFn(params ...string) (interface{}, error) {
	return Int(), nil
}

func int64InRangeFn(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Int64InRange(int64(min), int64(max)), nil
}

func int64Fn(params ...string) (interface{}, error) {
	return Int64(), nil
}

func int32InRangeFn(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Int32InRange(int32(min), int32(max)), nil
}

func int32Fn(params ...string) (interface{}, error) {
	return Int32(), nil
}

func int16InRangeFn(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Int16InRange(int16(min), int16(max)), nil
}

func int16Fn(params ...string) (interface{}, error) {
	return Int16(), nil
}

func int8InRangeFn(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Int8InRange(int8(min), int8(max)), nil
}

func int8Fn(params ...string) (interface{}, error) {
	return Int8(), nil
}

func uintInRangeFn(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return UintInRange(uint(min), uint(max)), nil
}

func uintFn(params ...string) (interface{}, error) {
	return Uint(), nil
}

func uint64InRangeFn(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Uint64InRange(uint64(min), uint64(max)), nil
}

func uint64Fn(params ...string) (interface{}, error) {
	return Uint64(), nil
}

func uint32InRangeFn(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Uint32InRange(uint32(min), uint32(max)), nil
}

func uint32Fn(params ...string) (interface{}, error) {
	return Uint32(), nil
}

func uint16InRangeFn(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Uint16InRange(uint16(min), uint16(max)), nil
}

func uint16Fn(params ...string) (interface{}, error) {
	return Uint16(), nil
}

func uint8InRangeFn(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Uint8InRange(uint8(min), uint8(max)), nil
}

func uint8Fn(params ...string) (interface{}, error) {
	return Uint8(), nil
}

func float64InRangeFn(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxFloat64(params...)
	if err != nil {
		return nil, err
	}
	return Float64InRange(min, max), nil
}

func float64Fn(params ...string) (interface{}, error) {
	return Float64(), nil
}

func float32InRangeFn(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxFloat64(params...)
	if err != nil {
		return nil, err
	}
	return Float32InRange(float32(min), float32(max)), nil
}

func float32Fn(params ...string) (interface{}, error) {
	return Float32(), nil
}
