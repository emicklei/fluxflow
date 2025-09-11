package internal

import "reflect"

var builtins = map[string]reflect.Value{}

func init() {
	builtins["int32"] = reflect.ValueOf(func(i int) int32 { return int32(i) })
}
