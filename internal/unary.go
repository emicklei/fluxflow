package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type UnaryExpr struct {
	*ast.UnaryExpr
	X Expr
}

func (u *UnaryExpr) String() string {
	return fmt.Sprintf("UnaryExpr(%s %s)", u.Op, u.X)
}

func (u *UnaryExpr) Eval(vm *VM) reflect.Value {
	return reflect.Value{}
}
