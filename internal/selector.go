package internal

import (
	"go/ast"
	"reflect"
)

type Selector struct {
	X Expr
	*ast.SelectorExpr
}

// return a function?
func (s Selector) Eval(env *Env) reflect.Value {
	/**
	look up X
	apply selector Sel
	**/
	_ = s.X.Eval(env)
	return reflect.Value{} // obj.Select(s.SelectorExpr.Sel.Name)
}
