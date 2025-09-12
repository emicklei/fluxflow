package internal

import "reflect"

type Env struct {
	parent      *Env
	symbolTable map[string]reflect.Value
}

func newEnv() *Env {
	return &Env{symbolTable: map[string]reflect.Value{}}
}
func (e *Env) subEnv() *Env {
	child := newEnv()
	child.parent = e
	return child
}
func (e *Env) lookUp(name string) reflect.Value {
	v, ok := e.symbolTable[name]
	if !ok {
		if e.parent == nil {
			return reflect.Value{}
		}
		return e.parent.lookUp(name)
	}
	return v
}

func (e *Env) ownerOf(name string) *Env {
	_, ok := e.symbolTable[name]
	if !ok {
		if e.parent == nil {
			return nil
		}
		return e.parent.ownerOf(name)
	}
	return e
}

func (e *Env) set(name string, value reflect.Value) {
	e.symbolTable[name] = value
}
