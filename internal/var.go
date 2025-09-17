package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

var _ CanAssign = ConstOrVar{}

type ConstOrVar struct {
	step
	spec *ast.ValueSpec
}

func (v ConstOrVar) declStep() Evaluable { return v }

func (v ConstOrVar) Assign(env *Env, value reflect.Value) {
	// TODO value->values?
	env.ownerOf(v.spec.Names[0].Name).set(v.spec.Names[0].Name, value)
}
func (v ConstOrVar) Define(env *Env, value reflect.Value) {
	// TODO value->values?
	env.set(v.spec.Names[0].Name, value)
}
func (v ConstOrVar) Eval(vm *VM) {
	vv := vm.localEnv().lookUp(v.spec.Names[0].Name)
	vm.Returns(vv)
}
func (v ConstOrVar) String() string {
	return fmt.Sprintf("Var(%v)", v.spec.Names)
}
