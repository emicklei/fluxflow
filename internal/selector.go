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
	recv := vm.ReturnsEval(s.X)
	rec, ok := recv.Interface().(FieldSelectable)
	if ok {
		sel := rec.Select(s.Sel.Name)
		vm.Returns(sel)
		return
	}
}

func (s SelectorExpr) Assign(env *Env, value reflect.Value) {
	//env.set(i.Obj.Name, value)
}

func (s SelectorExpr) String() string {
	return fmt.Sprintf("SelectorExpr(%v, %v)", s.X, s.SelectorExpr.Sel.Name)
}
