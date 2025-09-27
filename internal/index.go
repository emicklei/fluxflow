package internal

import (
	"fmt"
	"go/ast"
)

type IndexExpr struct {
	*ast.IndexExpr
	X     Expr
	Index Expr
}

func (i IndexExpr) String() string {
	return fmt.Sprintf("IndexExpr(%v, %v)", i.X, i.Index)
}
func (i IndexExpr) Eval(vm *VM) {}
