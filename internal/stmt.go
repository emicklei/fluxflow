package internal

import (
	"fmt"
	"go/ast"
)

type ExprStmt struct {
	*ast.ExprStmt
	X Expr
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

func (s DeclStmt) Eval(vm *VM) {
	s.Decl.declStep().Declare(vm)
}

func (s DeclStmt) String() string {
	return fmt.Sprintf("DeclStmt(%v)", s.Decl)
}

type Decl interface {
	declStep() CanDeclare
}

// LabeledStmt represents a labeled statement.
type LabeledStmt struct {
	*ast.LabeledStmt
	Label *Ident
	Stmt  Stmt
}

func (s LabeledStmt) String() string {
	return fmt.Sprintf("LabeledStmt(%v)", s.Label)
}

func (s LabeledStmt) stmtStep() Evaluable { return s }

func (s LabeledStmt) Eval(vm *VM) {
	s.Stmt.stmtStep().Eval(vm)
}

// BranchStmt represents a break, continue, goto, or fallthrough statement.
type BranchStmt struct {
	*ast.BranchStmt
	Label *Ident
}

func (s BranchStmt) Eval(vm *VM) {}

func (s BranchStmt) String() string {
	return fmt.Sprintf("BranchStmt(%v)", s.Label)
}

func (s BranchStmt) stmtStep() Evaluable { return s }
