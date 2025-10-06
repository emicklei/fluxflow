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

func (v ConstOrVar) Assign(vm *VM, value reflect.Value) {
	vm.localEnv().valueOwnerOf(v.Name.Name).set(v.Name.Name, value)
}
func (v ConstOrVar) Define(vm *VM, value reflect.Value) {
	vm.localEnv().set(v.Names[0].Name, value)
}
func (v ConstOrVar) Declare(vm *VM) bool {
	if v.Value != nil {
		actual := vm.returnsEval(v.Value)
		if !actual.IsValid() {
			return false
		}
		vm.localEnv().set(v.Name.Name, actual)
		return true
	}
	// if nil then zero
	if z, ok := v.Type.(HasZeroValue); ok {
		zv := z.ZeroValue(vm.localEnv())
		vm.localEnv().set(v.Name.Name, zv)
	}
	return true
}

func (v ConstOrVar) Eval(vm *VM) {
	vv := vm.localEnv().valueLookUp(v.Name.Name)
	vm.pushOperand(vv)
}
func (v ConstOrVar) String() string {
	return fmt.Sprintf("ConstOrVar(%v)", v.Name.Name)
}
