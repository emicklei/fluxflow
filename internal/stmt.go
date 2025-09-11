package internal

import (
	"go/ast"
)

type Stmt struct {
	step
	X Expr
	*ast.ExprStmt
}

func (s Stmt) Perform(vm *VM) {
	s.X.Eval(vm.env)
}
