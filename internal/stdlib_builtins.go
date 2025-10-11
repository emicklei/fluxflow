package internal

import (
	"fmt"
	"reflect"
)

// https://pkg.go.dev/builtin
var builtinsMap = map[string]reflect.Value{
	"int":        reflect.ValueOf(func(i int) int { return int(i) }),
	"uint":       reflect.ValueOf(func(i int) uint { return uint(i) }),
	"int8":       reflect.ValueOf(func(i int) int8 { return int8(i) }),
	"uint8":      reflect.ValueOf(func(i int) uint8 { return uint8(i) }),
	"byte":       reflect.ValueOf(func(i int) byte { return byte(i) }), // alias for uint8
	"int16":      reflect.ValueOf(func(i int) int16 { return int16(i) }),
	"uint16":     reflect.ValueOf(func(i int) uint16 { return uint16(i) }),
	"int32":      reflect.ValueOf(func(i int) int32 { return int32(i) }),
	"rune":       reflect.ValueOf(func(i int) rune { return rune(i) }), // alias for int32
	"uint32":     reflect.ValueOf(func(i int) uint32 { return uint32(i) }),
	"int64":      reflect.ValueOf(func(i int) int64 { return int64(i) }),
	"uint64":     reflect.ValueOf(func(i int) uint64 { return uint64(i) }),
	"uintptr":    reflect.ValueOf(func(i uint) uintptr { return uintptr(i) }),
	"float32":    reflect.ValueOf(func(f float64) float32 { return float32(f) }),
	"float64":    reflect.ValueOf(func(f float32) float64 { return float64(f) }),
	"complex64":  reflect.ValueOf(func(c complex128) complex64 { return complex64(c) }),
	"complex128": reflect.ValueOf(func(c complex64) complex128 { return complex128(c) }),
	"true":       reflect.ValueOf(true),  // not presented as Literal
	"false":      reflect.ValueOf(false), // not presented as Literal
	"print":      reflect.ValueOf(func(args ...any) { fmt.Print(args...) }),
	"println":    reflect.ValueOf(func(args ...any) { fmt.Println(args...) }),
}
