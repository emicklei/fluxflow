package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type CallExpr struct {
	Fun  Expr
	Args []Expr
	*ast.CallExpr
}

func (c CallExpr) Perform(vm *VM) {
	f := c.Fun.Eval(vm)
	args := make([]reflect.Value, len(c.Args))
	for i, arg := range c.Args {
		args[i] = arg.Eval(vm)
	}
	vals := f.Call(args)
	top := vm.callStack.pop()
	top.returnValues = vals
	vm.callStack.push(top)
}
func (c CallExpr) Eval(vm *VM) reflect.Value {
	f := c.Fun.Eval(vm)
	args := make([]reflect.Value, len(c.Args))
	for i, arg := range c.Args {
		args[i] = arg.Eval(vm)
	}
	vals := f.Call(args)
	if len(vals) == 0 {
		return reflect.Value{}
	}
	// tODO multiple return values
	return vals[0]
}

func (c CallExpr) Assign(env *Env, value reflect.Value) {}

func (c CallExpr) String() string {
	return fmt.Sprintf("CallExpr(%v, %d)", c.Fun, len(c.Args))
}
