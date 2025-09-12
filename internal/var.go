package internal

import (
	"go/ast"
	"reflect"
)

// For both variable and constant for now
type Var struct {
	step
	spec *ast.ValueSpec
}

func (v *Var) Assign(env *Env, value reflect.Value) {
	env.set(v.spec.Names[0].Name, value)
}
func (v *Var) Eval(env *Env) reflect.Value {
	return env.lookUp(v.spec.Names[0].Name)
}
