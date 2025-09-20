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

func (c CallExpr) evalAppend(vm *VM) {
	args := make([]reflect.Value, len(c.Args))
	for i, arg := range c.Args {
		args[i] = vm.ReturnsEval(arg)
	}
	result := reflect.Append(args[0], args[1:]...)
	vm.Returns(result)
}

func (c CallExpr) Eval(vm *VM) {
	if i, ok := c.Fun.(Ident); ok {
		if i.Name == "append" {
			c.evalAppend(vm)
			return
		}
	}
	// function f is either an external or an interpreted one
	f := vm.ReturnsEval(c.Fun)
	if f.Kind() == reflect.Func {

		args := make([]reflect.Value, len(c.Args))
		for i, arg := range c.Args {
			args[i] = vm.ReturnsEval(arg)
			//fmt.Printf("%d %v %v %T\n", i, arg, args[i], args[i])
		}
		vals := f.Call(args)

		// set return values on top of stack
		top := vm.callStack.pop()
		top.returnValues = vals
		vm.callStack.push(top)
	}
	if f.Kind() == reflect.Struct {
		lf, ok := f.Interface().(FuncDecl)
		if !ok {
			panic("expected FuncDecl, got " + fmt.Sprintf("%T", f.Interface()))
		}

		params := make([]reflect.Value, len(c.Args))
		for i, arg := range c.Args {
			params[i] = vm.ReturnsEval(arg)
		}
		fr := stackFrame{}
		fr.funcArgs = params
		fr.env = vm.env.subEnv()
		vm.callStack.push(fr)

		// take all parameters and put them in the env
		p := 0
		for _, field := range lf.Type.Params.List {
			for _, name := range field.Names {
				fr.env.set(name.Name, params[p])
				p++
			}
		}
		lf.Body.Eval(vm)
		vm.callStack.pop()
	}
}

func (c CallExpr) Assign(env *Env, value reflect.Value) {}

func (c CallExpr) String() string {
	return fmt.Sprintf("CallExpr(%v, len=%d)", c.Fun, len(c.Args))
}
