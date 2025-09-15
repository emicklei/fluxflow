package internal

import (
	"fmt"
	"go/ast"
)

type IncDecStmt struct {
	*ast.IncDecStmt
	X Expr
}

func (i *IncDecStmt) stmtStep() Evaluable { return i }

func (i *IncDecStmt) String() string {
	return fmt.Sprintf("IncDecStmt(%v)", i.X)
}
func (i *IncDecStmt) Eval(vm *VM) {
}
