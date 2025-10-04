package internal

import (
	"reflect"
)

// https://pkg.go.dev/builtin
var builtinsMap = map[string]reflect.Value{
	"int":     reflect.ValueOf(func(i int) int { return int(i) }),
	"uint":    reflect.ValueOf(func(i int) uint { return uint(i) }),
	"int8":    reflect.ValueOf(func(i int) int8 { return int8(i) }),
	"uint8":   reflect.ValueOf(func(i int) uint8 { return uint8(i) }),
	"int16":   reflect.ValueOf(func(i int) int16 { return int16(i) }),
	"uint16":  reflect.ValueOf(func(i int) uint16 { return uint16(i) }),
	"int32":   reflect.ValueOf(func(i int) int32 { return int32(i) }),
	"uint32":  reflect.ValueOf(func(i int) uint32 { return uint32(i) }),
	"int64":   reflect.ValueOf(func(i int) int64 { return int64(i) }),
	"uint64":  reflect.ValueOf(func(i int) uint64 { return uint64(i) }),
	"float32": reflect.ValueOf(func(f float64) float32 { return float32(f) }),
	"float64": reflect.ValueOf(func(f float32) float64 { return float64(f) }),
	"true":    reflect.ValueOf(true),  // not presented as Literal
	"false":   reflect.ValueOf(false), // not presented as Literal
}
