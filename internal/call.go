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

func (c CallExpr) Eval(vm *VM) reflect.Value {
	// function f is either an external or an interpreted one
	f := c.Fun.Eval(vm)
	if f.Kind() == reflect.Func {
		args := make([]reflect.Value, len(c.Args))
		for i, arg := range c.Args {
			args[i] = arg.Eval(vm)
		}
		vals := f.Call(args)
		top := vm.callStack.pop()
		top.returnValues = vals
		vm.callStack.push(top)
	}
	if f.Kind() == reflect.Pointer { // reflect pointer to a funcdecl
		lf := f.Interface().(*FuncDecl)
		return lf.Body.Eval(vm)
	}
	return reflect.Value{}
}

func (c CallExpr) Assign(env *Env, value reflect.Value) {}

func (c CallExpr) String() string {
	return fmt.Sprintf("CallExpr(%v, %d)", c.Fun, len(c.Args))
}
