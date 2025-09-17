package internal

import (
	"fmt"
	"go/ast"
)

type ExprStmt struct {
	X Expr
	*ast.ExprStmt
}

func (s ExprStmt) stmtStep() Evaluable { return s }

func (s ExprStmt) Eval(vm *VM) {
	s.X.Eval(vm)
}

func (s ExprStmt) String() string {
	return fmt.Sprintf("ExprStmt(%v)", s.X)
}

type DeclStmt struct {
	*ast.DeclStmt
	Decl Decl
}

func (s DeclStmt) stmtStep() Evaluable { return s }

func (s DeclStmt) Eval(vm *VM) {}

func (s DeclStmt) String() string {
	return fmt.Sprintf("DeclStmt(%v)", s.Decl)
}

type Decl interface {
	declStep() Evaluable
}
