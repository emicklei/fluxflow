package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type SelectorExpr struct {
	step
	X Expr
	*ast.SelectorExpr
}

// return a function?
func (s SelectorExpr) Eval(env *Env) reflect.Value {
	/**
	look up X
	apply selector Sel
	**/
	_ = s.X.Eval(env)
	return reflect.Value{} // obj.Select(s.SelectorExpr.Sel.Name)
}

func (s SelectorExpr) Assign(env *Env, value reflect.Value) {
	//env.set(i.Obj.Name, value)
}

func (s SelectorExpr) String() string {
	return fmt.Sprintf("SelectorExpr(%v, %v)", s.X, s.SelectorExpr.Sel.Name)
}
