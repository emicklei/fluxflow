package internal

import (
	"fmt"
	"go/ast"
	"go/token"
)

type AssignStmt struct {
	*ast.AssignStmt
	Lhs []Expr
	Rhs []Expr
}

func (a AssignStmt) stmtStep() Evaluable { return a }

func (a AssignStmt) Eval(vm *VM) {
	for _, each := range a.Rhs {
		// values are stacked operands
		each.Eval(vm)
	}
	// operands are stacked in reverse order
	for i := len(a.Lhs) - 1; i != -1; i-- {
		each := a.Lhs[i]
		v := vm.callStack.top().pop()
		target, ok_ := each.(CanAssign)
		if !ok_ {
			panic("cannot assign to " + fmt.Sprintf("%T", each))
		}
		switch a.AssignStmt.Tok {
		case token.DEFINE: // :=
			target.Define(vm, v)
		case token.ASSIGN: // =
			target.Assign(vm, v)
		default:
			panic("unsupported assignment " + a.AssignStmt.Tok.String())
		}
	}
}
func (a AssignStmt) String() string {
	return fmt.Sprintf("Assign(%v,%s, %v)", a.Lhs, a.AssignStmt.Tok, a.Rhs)
}
