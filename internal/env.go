package internal

import (
	"fmt"
	"reflect"
)

type Env interface {
	valueLookUp(name string) reflect.Value
	typeLookUp(name string) reflect.Type
	valueOwnerOf(name string) Env
	set(name string, value reflect.Value)
	newChildEnvironment() Env
	addConstOrVar(cv ConstOrVar)
}

type PkgEnvironment struct {
	Env
	pkgTable  map[string]ImportSpec
	declTable map[string]CanDeclare
}

func newPkgEnvironment(parent Env) Env {
	return &PkgEnvironment{
		Env:       newEnvironment(parent),
		pkgTable:  map[string]ImportSpec{},
		declTable: map[string]CanDeclare{},
	}
}

func (p *PkgEnvironment) addConstOrVar(cv ConstOrVar) {
	p.declTable[cv.Name.Name] = CanDeclare(cv)
}

type Environment struct {
	parent     Env
	valueTable map[string]reflect.Value
}

func newEnvironment(parentOrNil Env) Env {
	return &Environment{
		parent:     parentOrNil,
		valueTable: map[string]reflect.Value{},
	}
}
func (e *Environment) newChildEnvironment() Env {
	return newEnvironment(e)
}
func (e *Environment) valueLookUp(name string) reflect.Value {
	v, ok := e.valueTable[name]
	if !ok {
		if e.parent == nil {
			return reflect.Value{}
		}
		return e.parent.valueLookUp(name)
	}
	return v
}

func (e *Environment) typeLookUp(name string) reflect.Type {
	v, ok := builtinTypesMap[name]
	if !ok {
		return nil
	}
	return v
}

func (e *Environment) valueOwnerOf(name string) Env {
	_, ok := e.valueTable[name]
	if !ok {
		if e.parent == nil {
			return nil
		}
		return e.parent.valueOwnerOf(name)
	}
	return e
}

func (e *Environment) set(name string, value reflect.Value) {
	e.valueTable[name] = value
}

func (e *Environment) addConstOrVar(cv ConstOrVar) {
	//e.valueTable[cv.Name.Name] =
	fmt.Println(cv)
}
