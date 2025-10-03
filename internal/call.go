package internal

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/emicklei/structexplorer"
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
	// structexplorer.Break("vm", vm)
	if i, ok := c.Fun.(Ident); ok {
		if i.Name == "append" {
			c.evalAppend(vm)
			return
		}
	}
	// function f is either an external or an interpreted one
	f := vm.ReturnsEval(c.Fun)

	// todo: make a switch
	if f.Kind() == reflect.Func {

		args := make([]reflect.Value, len(c.Args))
		for i, arg := range c.Args {
			val := vm.ReturnsEval(arg)
			assertValid(c.String(), val)
			args[i] = val
		}
		vals := f.Call(args)

		// set return values on top of stack
		for _, each := range vals {
			vm.Returns(each)
		}
		return
		// top := vm.callStack.pop()
		// top.returnValues = vals
		// vm.callStack.push(top)
	}
	if f.Kind() == reflect.Struct {
		lf, ok := f.Interface().(FuncDecl)
		if !ok {
			panic("expected FuncDecl, got " + fmt.Sprintf("%T", f.Interface()))
		}

		structexplorer.Break("vm", vm)
		// prepare arguments
		args := make([]reflect.Value, len(c.Args))
		for i, arg := range c.Args {
			args[i] = vm.ReturnsEval(arg)
		}
		frame := vm.pushNewFrame()

		// take all parameters and put them in the env of the new frame
		p := 0
		for _, field := range lf.Type.Params.List {
			for _, name := range field.Names {
				frame.env.set(name.Name, args[p])
				p++
			}
		}
		lf.Body.Eval(vm)
		top := vm.popFrame()
		for _, each := range top.returnValues {
			vm.Returns(each)
		}
	}
}

func assertValid(role string, v reflect.Value) {
	if !v.IsValid() {
		panic(fmt.Sprintf("%s: invalid value", role))
	}
}

// used?
func (c CallExpr) Assign(env *Environment, value reflect.Value) {}

func (c CallExpr) String() string {
	return fmt.Sprintf("CallExpr(%v, len=%d)", c.Fun, len(c.Args))
}
