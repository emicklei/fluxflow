package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type ExprStmt struct {
	X Expr
	*ast.ExprStmt
}

func (s *ExprStmt) stmtStep() Evaluable { return s }

func (s ExprStmt) Eval(vm *VM) reflect.Value {
	return s.X.Eval(vm)
}

func (s ExprStmt) String() string {
	return fmt.Sprintf("ExprStmt(%v)", s.X)
}
