package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type Ident struct {
	step
	*ast.Ident
}

func (i Ident) Eval(env *Env) reflect.Value {
	return env.lookUp(i.Obj.Name)
}
func (i Ident) Assign(env *Env, value reflect.Value) {
	env.set(i.Obj.Name, value)
}
func (i Ident) String() string {
	if i.Obj == nil {
		return fmt.Sprintf("Ident(%v)", i.Name)
	}
	return fmt.Sprintf("Ident(%v)", i.Obj.Name)
}
