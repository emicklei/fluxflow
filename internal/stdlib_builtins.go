package internal

import "reflect"

// https://pkg.go.dev/builtin
var builtinsMap = map[string]reflect.Value{
	"int32": reflect.ValueOf(func(i int) int32 { return int32(i) }),
	"true":  reflect.ValueOf(true),
	"false": reflect.ValueOf(false),
	// "append": reflect.ValueOf(func[T any](slice []T, elems ...T) [] {
	// 	return append(slice, elems...)
	// }),
}
