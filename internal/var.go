package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

var _ CanAssign = ConstOrVar{}

type ConstOrVar struct {
	*ast.ValueSpec
	// for each Name in ValueSpec there is a ConstOrVar
	Name  *Ident
	Type  Expr
	Value Expr
}

func (v ConstOrVar) declStep() CanDeclare { return v }

func (v ConstOrVar) Assign(env Env, value reflect.Value) {
	env.valueOwnerOf(v.Name.Name).set(v.Name.Name, value)
}
func (v ConstOrVar) Define(env Env, value reflect.Value) {
	env.set(v.Names[0].Name, value)
}
func (v ConstOrVar) Declare(vm *VM) {
	if v.Value != nil {
		vm.localEnv().set(v.Name.Name, vm.ReturnsEval(v.Value))
		return
	}
	// if nil then zero
	if z, ok := v.Type.(HasZeroValue); ok {
		zv := z.ZeroValue(vm.localEnv())
		vm.localEnv().set(v.Name.Name, zv)
	}
}

func (v ConstOrVar) Eval(vm *VM) {
	vv := vm.localEnv().valueLookUp(v.Name.Name)
	vm.Returns(vv)
}
func (v ConstOrVar) String() string {
	return fmt.Sprintf("ConstOrVar(%v)", v.Name.Name)
}
