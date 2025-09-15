package internal

import (
	"fmt"
	"go/ast"
)

type ExprStmt struct {
	X Expr
	*ast.ExprStmt
}

func (s *ExprStmt) stmtStep() Evaluable { return s }

func (s ExprStmt) Eval(vm *VM) {
	s.X.Eval(vm)
}

func (s ExprStmt) String() string {
	return fmt.Sprintf("ExprStmt(%v)", s.X)
}
