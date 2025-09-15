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

func (c CallExpr) Eval(vm *VM) {
	// function f is either an external or an interpreted one
	f := vm.ReturnsEval(c.Fun)
	if f.Kind() == reflect.Func {

		fr := stackFrame{}
		fr.env = vm.env.subEnv()

		args := make([]reflect.Value, len(c.Args))
		for i, arg := range c.Args {
			args[i] = vm.ReturnsEval(arg)
		}

		vm.callStack.push(fr)
		vals := f.Call(args)

		// replace with returns
		top := vm.callStack.pop()
		top.returnValues = vals
		vm.callStack.push(top)
	}
	if f.Kind() == reflect.Pointer { // reflect pointer to a funcdecl
		lf := f.Interface().(*FuncDecl)

		fr := stackFrame{}
		fr.env = vm.env.subEnv()
		vm.callStack.push(fr)

		lf.Body.Eval(vm)
	}
}

func (c CallExpr) Assign(env *Env, value reflect.Value) {}

func (c CallExpr) String() string {
	return fmt.Sprintf("CallExpr(%v, %d)", c.Fun, len(c.Args))
}
