package internal

import (
	"go/ast"
	"reflect"
)

type Var struct {
	step
	spec *ast.ValueSpec
}

func (v *Var) declStep() Evaluable { return v }

func (v *Var) Assign(env *Env, value reflect.Value) {
	// TODO value->values?
	env.set(v.spec.Names[0].Name, value)
}
func (v *Var) Eval(vm *VM) {
	vv := vm.localEnv().lookUp(v.spec.Names[0].Name)
	vm.Returns(vv)
}
