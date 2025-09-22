package internal

import (
	"fmt"
	"go/ast"
)

type KeyValueExpr struct {
	*ast.KeyValueExpr
	Key   Expr
	Value Expr
}

func (e KeyValueExpr) String() string {
	return fmt.Sprintf("KeyValueExpr(%v,%v)", e.Key, e.Value)
}

func (e KeyValueExpr) Eval(vm *VM) {

}
