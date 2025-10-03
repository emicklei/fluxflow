package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

var _ CanAssign = IndexExpr{}

type IndexExpr struct {
	*ast.IndexExpr
	X     Expr
	Index Expr
}

func (i IndexExpr) String() string {
	return fmt.Sprintf("IndexExpr(%v, %v)", i.X, i.Index)
}
func (i IndexExpr) Eval(vm *VM) {
	target := vm.ReturnsEval(i.X)
	index := vm.ReturnsEval(i.Index)
	if target.Kind() == reflect.Map {
		vm.Returns(target.MapIndex(index))
		return
	}
	if target.Kind() == reflect.Slice || target.Kind() == reflect.Array {
		vm.Returns(target.Index(int(index.Int())))
		return
	}
	expected(target, "map or slice or array")
}

func (i IndexExpr) Assign(vm *VM, value reflect.Value) {
	target := vm.ReturnsEval(i.X)
	index := vm.ReturnsEval(i.Index)
	if target.Kind() == reflect.Map {
		target.SetMapIndex(index, value)
		return
	}
	if target.Kind() == reflect.Slice || target.Kind() == reflect.Array {
		reflect.ValueOf(target.Interface()).Index(int(index.Int())).Set(value)
		return
	}
	expected(target, "map or slice or array")
}

func (i IndexExpr) Define(vm *VM, value reflect.Value) {
	fmt.Println("IndexExpr.Define", i, value)
}
