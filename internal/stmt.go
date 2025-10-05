package internal

import (
	"fmt"
	"go/ast"
	"reflect"
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

// A SwitchStmt represents an expression switch statement.
type SwitchStmt struct {
	*ast.SwitchStmt
	Init Stmt // initialization statement; or nil
	Tag  Expr // tag expression; or nil
	Body BlockStmt
}

func (s SwitchStmt) stmtStep() Evaluable { return s }

func (s SwitchStmt) Eval(vm *VM) {
	vm.pushNewFrame()
	defer vm.popFrame() // to handle break statements
	if s.Init != nil {
		s.Init.stmtStep().Eval(vm)
	}
	if s.Tag != nil {
		s.Tag.Eval(vm)
	}
	s.Body.Eval(vm)
}
func (s SwitchStmt) String() string {
	return fmt.Sprintf("SwitchStmt(%v,%v,%v)", s.Init, s.Tag, s.Body)
}

// A CaseClause represents a case of an expression or type switch statement.
type CaseClause struct {
	*ast.CaseClause
	List []Expr // list of expressions; nil means default case
	Body []Stmt
}

func (c CaseClause) String() string {
	return fmt.Sprintf("CaseClause(%v,%v)", c.List, c.Body)
}
func (c CaseClause) Eval(vm *VM) {
	if c.List == nil {
		// default case
		for _, stmt := range c.Body {
			stmt.stmtStep().Eval(vm)
		}
		return
	}
	f := vm.callStack.top()
	var left reflect.Value
	if !f.operandStack.isEmpty() {
		left = vm.callStack.top().pop()
	}
	for _, expr := range c.List {
		right := vm.ReturnsEval(expr)
		var cond bool
		if left.IsValid() {
			// because value is on the operand stack we compare
			cond = left.Equal(right)
		} else {
			// no operand on stack, treat as boolean expression
			cond = right.Bool()
		}
		if cond {
			vm.pushNewFrame()
			defer vm.popFrame()
			for _, stmt := range c.Body {
				stmt.stmtStep().Eval(vm)
			}
			return
		}
	}
}

func (c CaseClause) stmtStep() Evaluable { return c }
