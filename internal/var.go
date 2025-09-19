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
func (v ConstOrVar) Declare(vm *VM) {
	if len(v.Values) > 0 {
		for i, val := range v.Values {
			vm.env.set(v.spec.Names[i].Name, vm.ReturnsEval(val))
		}
		return
	}
	if z, ok := v.Type.(HasZeroValue); ok {
		zv := z.ZeroValue(vm.env)
		for _, name := range v.spec.Names {
			vm.env.set(name.Name, zv)
		}
	}
}

func (v ConstOrVar) Eval(vm *VM) {
	vv := vm.localEnv().valueLookUp(v.spec.Names[0].Name)
	vm.Returns(vv)
}
func (v ConstOrVar) String() string {
	return fmt.Sprintf("Var(%v)", v.spec.Names)
}
