package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type StarExpr struct {
	X Expr
	*ast.StarExpr
}

func (s StarExpr) Eval(vm *VM) reflect.Value {
	v := s.X.Eval(vm)
	return v.Elem()
}
func (s StarExpr) Assign(vm *VM, value reflect.Value) {
	v := s.X.Eval(vm)
	v.Elem().Set(value)
}
func (s StarExpr) String() string {
	return fmt.Sprintf("StarExpr(%v)", s.X)
}
