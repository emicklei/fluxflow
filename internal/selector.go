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
	if !recv.IsValid() {
		panic("not a valid receiver")
	}
	rec, ok := recv.Interface().(FieldSelectable)
	if ok {
		sel := rec.Select(s.Sel.Name)
		vm.Returns(sel)
		return
	}
	panic("not implemented: not FieldSelectable" + recv.String())
}

func (s SelectorExpr) Assign(env *Env, value reflect.Value) {
	println("not implemented: SelectorExpr:Assign")
	//env.set(i.Obj.Name, value)
}

func (s SelectorExpr) String() string {
	return fmt.Sprintf("SelectorExpr(%v, %v)", s.X, s.SelectorExpr.Sel.Name)
}
