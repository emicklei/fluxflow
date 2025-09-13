package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type CallExpr struct {
	step
	Fun  Expr
	Args []Expr
	*ast.CallExpr
}

func (c CallExpr) Perform(vm *VM) {
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
func (c CallExpr) Eval(env *Env) reflect.Value { return reflect.Value{} }

func (c CallExpr) Assign(env *Env, value reflect.Value) {}

func (c CallExpr) String() string {
	return fmt.Sprintf("CallExpr(%v, %d)", c.Fun, len(c.Args))
}
