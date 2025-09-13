package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type AssignStmt struct {
	Lhs []Expr
	Rhs []Expr
	*ast.AssignStmt
}

func (a AssignStmt) stmtStep() Evaluable { return a }

func (a AssignStmt) Eval(vm *VM) reflect.Value {
	for i, rhs := range a.Rhs {
		v := rhs.Eval(vm)
		target, ok_ := a.Lhs[i].(CanAssign)
		if !ok_ {
			panic("cannot assign to " + fmt.Sprintf("%T", a.Lhs[i]))
		}
		target.Assign(vm.localEnv(), v)
	}
	return reflect.Value{}
}

func (a AssignStmt) String() string {
	return fmt.Sprintf("Assign(%v %s)", a.Lhs, a.AssignStmt.Tok)
}
