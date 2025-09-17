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
	for i, rhs := range a.Rhs {
		v := vm.ReturnsEval(rhs)
		target, ok_ := a.Lhs[i].(CanAssign)
		if !ok_ {
			panic("cannot assign to " + fmt.Sprintf("%T", a.Lhs[i]))
		}
		switch a.AssignStmt.Tok {
		case token.DEFINE: // :=
			target.Define(vm.localEnv(), v)
		case token.ASSIGN: // =
			target.Assign(vm.localEnv(), v)
		default:
			panic("unsupported assignment " + a.AssignStmt.Tok.String())
		}
	}
}
func (a AssignStmt) String() string {
	return fmt.Sprintf("Assign(%v %s)", a.Lhs, a.AssignStmt.Tok)
}
