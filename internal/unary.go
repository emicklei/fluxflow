package internal

import (
	"fmt"
	"go/ast"
)

type UnaryExpr struct {
	*ast.UnaryExpr
	X Expr
}

func (u *UnaryExpr) String() string {
	return fmt.Sprintf("UnaryExpr(%s %s)", u.Op, u.X)
}

func (u *UnaryExpr) Eval(vm *VM) {
	u.X.Eval(vm)
}
