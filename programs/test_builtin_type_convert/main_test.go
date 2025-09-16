package main

import (
	"reflect"
	"testing"
)

func TestTypeConvert(t *testing.T) {
	var i32 int32
	var a16 int16 = 1
	ra16 := reflect.ValueOf(a16)
	a32 := ra16.Convert(reflect.TypeOf(i32))
	t.Logf("%T %v", a32.Interface(), a32.Interface())
}
