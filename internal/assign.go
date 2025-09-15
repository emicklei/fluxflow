package internal

import (
	"fmt"
	"go/ast"
)

type AssignStmt struct {
	Lhs []Expr
	Rhs []Expr
	*ast.AssignStmt
}

func (a AssignStmt) stmtStep() Evaluable { return a }

func (a AssignStmt) Eval(vm *VM) {
	for i, rhs := range a.Rhs {
		v := vm.ReturnsEval(rhs)
		target, ok_ := a.Lhs[i].(CanAssign)
		if !ok_ {
			panic("cannot assign to " + fmt.Sprintf("%T", a.Lhs[i]))
		}
		target.Assign(vm.localEnv(), v)
	}
}
func (a AssignStmt) String() string {
	return fmt.Sprintf("Assign(%v %s)", a.Lhs, a.AssignStmt.Tok)
}
