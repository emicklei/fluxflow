package internal

import "reflect"

// used?
type environment interface {
	valueLookUp(name string) reflect.Value
	typeLookUp(name string) reflect.Type
	valueOwnerOf(name string) environment
	set(name string, value reflect.Value)
}

type Env struct {
	parent     *Env
	valueTable map[string]reflect.Value
	pkgTable   map[string]ImportSpec
	declTable  map[string]Decl
}

func newEnv() *Env {
	return &Env{
		valueTable: map[string]reflect.Value{},
		pkgTable:   map[string]ImportSpec{},
		declTable:  map[string]Decl{},
	}
}
func (e *Env) subEnv() *Env {
	child := newEnv()
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

func (e *Env) valueOwnerOf(name string) *Env {
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

func (e *Env) setDecl(name string, decl Decl) {
	e.declTable[name] = decl
}
