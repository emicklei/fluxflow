package internal

import "reflect"

type Env struct {
	parent     *Env
	valueTable map[string]reflect.Value
}

func newEnv() *Env {
	return &Env{valueTable: map[string]reflect.Value{}}
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
