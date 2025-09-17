package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

var _ CanAssign = new(Ident)

type Ident struct {
	step
	*ast.Ident
}

func (i Ident) Eval(vm *VM) {
	name := i.Name
	// TODO why?
	if i.Obj != nil {
		name = i.Obj.Name
	}
	vm.Returns(vm.localEnv().lookUp(name))
}
func (i Ident) Assign(env *Env, value reflect.Value) {
	env.ownerOf(i.Name).set(i.Name, value)
}
func (i Ident) Define(env *Env, value reflect.Value) {
	env.set(i.Name, value)
}
func (i Ident) String() string {
	if i.Obj == nil {
		return fmt.Sprintf("Ident(%v)", i.Name)
	}
	return fmt.Sprintf("Ident(%v)", i.Obj.Name)
}
