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
	// structexplorer.Break("vm", vm)
	if i, ok := c.Fun.(Ident); ok {
		if i.Name == "append" {
			c.evalAppend(vm)
			return
		}
	}
	// function f is either an external or an interpreted one
	f := vm.ReturnsEval(c.Fun)

	if !f.IsValid() {
		vm.doPanic("call to invalid function")
	}

	switch f.Kind() {
	case reflect.Func:
		args := make([]reflect.Value, len(c.Args))
		for i, arg := range c.Args {
			val := vm.ReturnsEval(arg)
			args[i] = val
		}
		vals := f.Call(args)
		// set return values on top of stack
		for _, each := range vals {
			vm.Returns(each)
		}

	case reflect.Struct:
		fl, ok := f.Interface().(FuncLit)
		if ok {
			c.handleFuncLit(vm, fl)
			return
		}
		lf, ok := f.Interface().(FuncDecl)
		if ok {
			c.handleFuncDecl(vm, lf)
			return
		}
		vm.doPanic("expected FuncDecl, got " + fmt.Sprintf("%T", f.Interface()))
	default:
		vm.doPanic("call to unknown function type")
	}
}

func (c CallExpr) handleFuncLit(vm *VM, fl FuncLit) {
	// TODO deduplicate with handleFuncDecl
	// prepare arguments
	args := make([]reflect.Value, len(c.Args))
	for i, arg := range c.Args {
		val := vm.ReturnsEval(arg)
		args[i] = val
	}
	frame := vm.pushNewFrame()
	// take all parameters and put them in the env of the new frame
	p := 0
	for _, field := range fl.Type.Params.List {
		for _, name := range field.Names {
			frame.env.set(name.Name, args[p])
			p++
		}
	}
	fl.Body.Eval(vm)
	top := vm.popFrame()
	for _, each := range top.returnValues {
		vm.Returns(each)
	}
}
func (c CallExpr) handleFuncDecl(vm *VM, fd FuncDecl) {
	// TODO deduplicate with handleFuncLit
	// prepare arguments
	args := make([]reflect.Value, len(c.Args))
	for i, arg := range c.Args {
		val := vm.ReturnsEval(arg)
		args[i] = val
	}
	frame := vm.pushNewFrame()
	// take all parameters and put them in the env of the new frame
	p := 0
	for _, field := range fd.Type.Params.List {
		for _, name := range field.Names {
			frame.env.set(name.Name, args[p])
			p++
		}
	}
	fd.Body.Eval(vm)
	top := vm.popFrame()
	for _, each := range top.returnValues {
		vm.Returns(each)
	}
}

// used?
func (c CallExpr) Assign(env *Environment, value reflect.Value) {}

func (c CallExpr) String() string {
	return fmt.Sprintf("CallExpr(%v, len=%d)", c.Fun, len(c.Args))
}
