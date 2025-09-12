package internal

import (
	"go/ast"
	"reflect"
)

type Ident struct {
	operatorUnimplemented
	step
	*ast.Ident
}

func (i Ident) Eval(env *Env) reflect.Value {
	return env.lookUp(i.Obj.Name)
}
