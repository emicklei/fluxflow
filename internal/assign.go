package internal

import (
	"go/ast"
	"reflect"
)

type Assign struct {
	step
	Lhs []Expr
	Rhs []Expr
	*ast.AssignStmt
}

func (a Assign) Eval(env *Env) reflect.Value {
	for i, rhs := range a.Rhs {
		v := rhs.Eval(env)
		a.Lhs[i].Assign(env, v)
	}
	return reflect.Value{}
}
