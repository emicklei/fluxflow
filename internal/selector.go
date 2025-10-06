package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type SelectorExpr struct {
	*ast.SelectorExpr
	X Expr
}

func (s SelectorExpr) Eval(vm *VM) {
	recv := vm.returnsEval(s.X)
	if !recv.IsValid() {
		panic("not a valid receiver")
	}
	rec, ok := recv.Interface().(FieldSelectable)
	if ok {
		sel := rec.Select(s.Sel.Name)
		vm.pushOperand(sel)
		return
	}
	panic("expected FieldSelectable: " + recv.String())
}

func (s SelectorExpr) Assign(env *Environment, value reflect.Value) {
	panic("not implemented: SelectorExpr:Assign")
	//env.set(i.Obj.Name, value)
}

func (s SelectorExpr) String() string {
	return fmt.Sprintf("SelectorExpr(%v, %v)", s.X, s.SelectorExpr.Sel.Name)
}
