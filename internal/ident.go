package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type Ident struct {
	step
	*ast.Ident
}

func (i Ident) Eval(vm *VM) reflect.Value {
	name := i.Name
	// TODO why?
	if i.Obj != nil {
		name = i.Obj.Name
	}
	return vm.localEnv().lookUp(name)
}
func (i Ident) Assign(vm *VM, value reflect.Value) {
	name := i.Name
	// TODO why?
	if i.Obj != nil {
		name = i.Obj.Name
	}
	vm.localEnv().set(name, value)
}
func (i Ident) String() string {
	if i.Obj == nil {
		return fmt.Sprintf("Ident(%v)", i.Name)
	}
	return fmt.Sprintf("Ident(%v)", i.Obj.Name)
}
