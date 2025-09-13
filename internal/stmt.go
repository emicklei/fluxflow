package internal

import (
	"fmt"
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
func (s Stmt) String() string {
	return fmt.Sprintf("Stmt(%v)", s.X)
}
