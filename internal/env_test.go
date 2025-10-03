package internal

import (
	"reflect"
	"testing"
)

func TestPackageEnv(t *testing.T) {
	global := newEnvironment(nil)
	global.set("a", reflect.ValueOf(1))
	pkg := newPkgEnvironment(global)
	fn := pkg.newChild()
	one := fn.valueLookUp("a")
	if one.Interface() != 1 {
		t.Fail()
	}
}
