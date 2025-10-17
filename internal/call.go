package internal

import (
	"fmt"
	"go/ast"
	"reflect"
	"slices"
)

var _ Flowable = CallExpr{}
var _ Expr = CallExpr{}

type CallExpr struct {
	*ast.CallExpr
	Fun  Expr
	Args []Expr
}

func (c CallExpr) Eval(vm *VM) {
	if i, ok := c.Fun.(Ident); ok {
		switch i.Name {
		case "delete":
			c.evalDelete(vm)
			return
		case "append":
			c.evalAppend(vm)
			return
		case "clear":
			cleared := c.evalClear(vm)
			// the argument of clear needs to be replaced
			if identArg, ok := c.Args[0].(Ident); ok {
				vm.callStack.top().env.set(identArg.Name, cleared)
			}
			return
		case "min":
			c.evalMin(vm)
			return
		case "max":
			c.evalMax(vm)
			return
		case "make":
			c.evalMake(vm)
			return
		}
	}
	// function f is either an external or an interpreted one
	f := vm.returnsEval(c.Fun)

	if !f.IsValid() {
		vm.fatal("call to invalid function:" + fmt.Sprintf("%v", c.Fun))
	}

	switch f.Kind() {
	case reflect.Func:
		args := make([]reflect.Value, len(c.Args))
		for i, arg := range c.Args {
			var val reflect.Value
			if vm.isStepping {
				val = vm.callStack.top().pop() // first to last, see Flow
			} else {
				val = vm.returnsEval(arg)
			}
			if !val.IsValid() {
				vm.fatal("call to function with invalid argument:" + fmt.Sprintf("%d=%v", i, arg))
			}
			args[i] = val
		}
		vals := f.Call(args)
		// set return values on top of stack
		for _, each := range vals {
			vm.pushOperand(each)
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
		vm.fatal("expected FuncDecl, got " + fmt.Sprintf("%T", f.Interface()))
	default:
		vm.fatal("call to unknown function type")
	}
}

func (c CallExpr) handleFuncLit(vm *VM, fl FuncLit) {
	// TODO deduplicate with handleFuncDecl
	// prepare arguments
	args := make([]reflect.Value, len(c.Args))
	for i, arg := range c.Args {
		var val reflect.Value
		if vm.isStepping {
			val = vm.callStack.top().pop() // first to last, see Flow
		} else {
			val = vm.returnsEval(arg)
		}
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
	if vm.isStepping {
		// when stepping we already have the call graph in FuncLit
		vm.takeAll(fl.callGraph)
	} else {
		vm.eval(fl.Body)
	}
	top := vm.popFrame()
	for _, each := range top.returnValues {
		vm.pushOperand(each)
	}
}
func (c CallExpr) handleFuncDecl(vm *VM, fd FuncDecl) {
	// TODO deduplicate with handleFuncLit
	// prepare arguments
	args := make([]reflect.Value, len(c.Args))
	for i, arg := range c.Args {
		var val reflect.Value
		if vm.isStepping {
			val = vm.callStack.top().pop() // first to last, see Flow
		} else {
			val = vm.returnsEval(arg)
		}
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
	if vm.isStepping {
		// when stepping we already have the call graph in FuncDecl
		vm.takeAll(fd.callGraph)
	} else {
		vm.eval(fd.Body)
	}
	top := vm.popFrame()
	// TODO optimize this
	slices.Reverse(top.returnValues)
	for _, each := range top.returnValues {
		vm.pushOperand(each)
	}
}

func (c CallExpr) Flow(g *grapher) (head Step) {
	// make sure first value is on top of the operand stack
	// so we can pop in the right order during Eval
	for i := len(c.Args) - 1; i >= 0; i-- {
		if i == len(c.Args)-1 {
			head = c.Args[i].Flow(g)
			continue
		}
		c.Args[i].Flow(g)
	}
	g.next(c)
	if head == nil {
		head = g.current
	}
	return head
}

func (c CallExpr) String() string {
	return fmt.Sprintf("CallExpr(%v, len=%d)", c.Fun, len(c.Args))
}
