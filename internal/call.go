package internal

import (
	"go/ast"
	"reflect"
)

type Call struct {
	Step
	Fun  Ident
	Args []Expr
	*ast.CallExpr
}

func (c Call) Perform(vm *VM) {
	f := c.Fun.Eval(vm.env)
	args := make([]reflect.Value, len(c.Args))
	for i, arg := range c.Args {
		args[i] = arg.Eval(vm.env)
	}
	vals := f.Call(args)
	top := vm.callStack.pop()
	top.returnValues = vals
	vm.callStack.push(top)
}
func (c Call) Eval(env *Env) reflect.Value { return reflect.Value{} }
