package internal

import (
	"reflect"
)

// https://pkg.go.dev/builtin
var builtinsMap = map[string]reflect.Value{
	"int8":    reflect.ValueOf(func(i int) int8 { return int8(i) }),
	"int16":   reflect.ValueOf(func(i int) int16 { return int16(i) }),
	"int32":   reflect.ValueOf(func(i int) int32 { return int32(i) }),
	"int64":   reflect.ValueOf(func(i int) int64 { return int64(i) }),
	"float32": reflect.ValueOf(func(f float64) float32 { return float32(f) }),
	"float64": reflect.ValueOf(func(f float32) float64 { return float64(f) }),
	"true":    reflect.ValueOf(true),  // not presented as Literal
	"false":   reflect.ValueOf(false), // not presented as Literal
}
