package data

type Pool []interface{}
type PoolGroup map[string]Pool
type PoolData map[string]PoolGroup

var DB = PoolData{
	"address":  Address,
	"country":  Country,
	"currency": Currency,
	"gender":   Gender,
	"lang":     Lang,
	"name":     Name,
	"timezone": Timezone,
}
