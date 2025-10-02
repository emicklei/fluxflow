package internal

import (
	"fmt"
	"reflect"
)

type ienv interface {
	valueLookUp(name string) reflect.Value
	typeLookUp(name string) reflect.Type
	valueOwnerOf(name string) ienv
	set(name string, value reflect.Value)
	subEnv() ienv
	addConstOrVar(cv ConstOrVar)
}

type PackageEnv struct {
	*Env
	pkgTable  map[string]ImportSpec
	declTable map[string]CanDeclare
}

func newPackageEnv(parent ienv) ienv {
	env := newEnv().(*Env)
	env.parent = parent
	return &PackageEnv{
		Env:       env,
		pkgTable:  map[string]ImportSpec{},
		declTable: map[string]CanDeclare{},
	}
}

func (p *PackageEnv) addConstOrVar(cv ConstOrVar) {
	p.declTable[cv.Name.Name] = CanDeclare(cv)
}

type Env struct {
	parent     ienv
	valueTable map[string]reflect.Value
	pkgTable   map[string]ImportSpec // only relevant on package level
}

func newEnv() ienv {
	return &Env{
		valueTable: map[string]reflect.Value{},
		pkgTable:   map[string]ImportSpec{},
	}
}
func (e *Env) subEnv() ienv {
	child := newEnv().(*Env)
	child.parent = e
	return child
}
func (e *Env) valueLookUp(name string) reflect.Value {
	v, ok := e.valueTable[name]
	if !ok {
		if e.parent == nil {
			return reflect.Value{}
		}
		return e.parent.valueLookUp(name)
	}
	return v
}

func (e *Env) typeLookUp(name string) reflect.Type {
	v, ok := builtinTypesMap[name]
	if !ok {
		return nil
	}
	return v
}

func (e *Env) valueOwnerOf(name string) ienv {
	_, ok := e.valueTable[name]
	if !ok {
		if e.parent == nil {
			return nil
		}
		return e.parent.valueOwnerOf(name)
	}
	return e
}

func (e *Env) set(name string, value reflect.Value) {
	e.valueTable[name] = value
}

func (e *Env) addConstOrVar(cv ConstOrVar) {
	//e.valueTable[cv.Name.Name] =
	fmt.Println(cv)
}
