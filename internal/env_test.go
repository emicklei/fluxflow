package internal

import (
	"reflect"
	"testing"
)

func TestPackageEnv(t *testing.T) {
	global := newEnv()
	global.set("a", reflect.ValueOf(1))
	pkg := newPackageEnv(global)
	fn := pkg.subEnv()
	one := fn.valueLookUp("a")
	if one.Interface() != 1 {
		t.Fail()
	}
}
