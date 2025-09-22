package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

var _ CanAssign = Ident{}

type Ident struct {
	*ast.Ident
}

func (i Ident) Eval(vm *VM) {
	vm.Returns(vm.localEnv().valueLookUp(i.Name))
}
func (i Ident) Assign(env *Env, value reflect.Value) {
	env.valueOwnerOf(i.Name).set(i.Name, value)
}
func (i Ident) Define(env *Env, value reflect.Value) {
	env.set(i.Name, value)
}

// ZeroValue returns the zero value iff the Ident represents a standard type.
func (i Ident) ZeroValue(env *Env) reflect.Value {
	// TODO handle interpreted types.
	rt := env.typeLookUp(i.Name)
	if rt == nil { // invalid
		return reflect.Value{}
	}
	return reflect.Zero(rt)
}

func (i Ident) String() string {
	if i.Obj == nil {
		return fmt.Sprintf("Ident(%v)", i.Name)
	}
	return fmt.Sprintf("Ident(%v)", i.Obj.Name)
}

// The identifier refers to a composeable type
func (i Ident) LiteralCompose(composite reflect.Value, values ...reflect.Value) reflect.Value {
	return reflect.Value{}
}
