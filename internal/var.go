package internal

import (
	"go/ast"
	"reflect"
)

// For both variable and constant for now
type Var struct {
	step
	spec *ast.ValueSpec
}

func (v *Var) Assign(env *Env, value reflect.Value) {
	env.set(v.spec.Names[0].Name, value)
}
func (v *Var) Eval(vm *VM) {
	vv := vm.localEnv().lookUp(v.spec.Names[0].Name)
	vm.Returns(vv)
}
