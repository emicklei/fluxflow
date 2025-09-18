package internal

import (
	"reflect"
	"testing"
)

func TestReflectIntSlice(t *testing.T) {
	rt := builtinTypesMap["int"]
	st := reflect.SliceOf(rt)
	rs := reflect.MakeSlice(st, 0, 0)
	rs = reflect.Append(rs, reflect.ValueOf(1))
	t.Log(rs)
}
