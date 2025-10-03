package internal

import (
	"fmt"
	"reflect"
)

func expected(value any, expectation string) reflect.Value {
	panic(fmt.Sprintf("expected %s : %v (%T)", expectation, value, value))
}

func mustString(v reflect.Value) string {
	if !v.IsValid() {
		panic("value not valid")
	}
	if !v.CanInterface() {
		panic("cannot get interface")
	}
	s, ok := v.Interface().(string)
	if !ok {
		panic(fmt.Sprintf("expected string but got %T", v.Interface()))
	}
	return s
}
