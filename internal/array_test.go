package internal

import (
	"reflect"
	"testing"
)

func TestReflectIntSlice(t *testing.T) {
	ra := []int{}
	rt := reflect.TypeOf(ra)
	rs := reflect.MakeSlice(rt, 0, 0)
	rs = reflect.Append(rs, reflect.ValueOf(1))
	t.Log(rs)
}
