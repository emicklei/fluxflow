package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

var _ CanAssign = ConstOrVar{}

type ConstOrVar struct {
	spec   *ast.ValueSpec
	Type   Expr
	Values []Expr
}

func (v ConstOrVar) declStep() CanDeclare { return v }

func (v ConstOrVar) Assign(env *Env, value reflect.Value) {
	// TODO value->values?
	env.valueOwnerOf(v.spec.Names[0].Name).set(v.spec.Names[0].Name, value)
}
func (v ConstOrVar) Define(env *Env, value reflect.Value) {
	// TODO value->values?
	env.set(v.spec.Names[0].Name, value)
}
func (v ConstOrVar) Declare(env *Env) {
	// TODO value->values?
	if z, ok := v.Type.(HasZeroValue); ok {
		env.set(v.spec.Names[0].Name, z.ZeroValue(env))
	}
}

func (v ConstOrVar) Eval(vm *VM) {
	vv := vm.localEnv().valueLookUp(v.spec.Names[0].Name)
	vm.Returns(vv)
}
func (v ConstOrVar) String() string {
	return fmt.Sprintf("Var(%v)", v.spec.Names)
}
