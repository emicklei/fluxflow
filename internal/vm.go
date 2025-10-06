package internal

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/emicklei/structexplorer"
)

type stackFrame struct {
	env          Env
	operandStack stack[reflect.Value]
	returnValues []reflect.Value
}

// push adds a value onto the operand stack.
func (f *stackFrame) push(v reflect.Value) {
	f.operandStack.push(v)
}

func (f stackFrame) pushOperand(v reflect.Value) stackFrame {
	f.operandStack.push(v)
	return f
}
func (f stackFrame) pushReturnValue(v reflect.Value) stackFrame {
	f.returnValues = append(f.returnValues, v)
	return f
}

// pop removes and returns the top value from the operand stack.
func (f *stackFrame) pop() reflect.Value {
	return f.operandStack.pop()
}

type VM struct {
	callStack stack[*stackFrame] // TODO use value io pointer?
	output    *bytes.Buffer
}

func newVM(env Env) *VM {
	vm := &VM{output: new(bytes.Buffer)}
	frame := &stackFrame{env: env}
	vm.callStack.push(frame)
	return vm
}

func (vm *VM) localEnv() Env {
	return vm.callStack.top().env
}

// ReturnsEval evaluates the argument and returns its return value.
func (vm *VM) ReturnsEval(e Evaluable) reflect.Value {
	if trace {
		fmt.Println("VM.ReturnsEval", e)
	}
	vm.eval(e)
	return vm.callStack.top().pop()
}

// Returns pushes a value onto the operand stack as the result of an evaluation.
func (vm *VM) Returns(v reflect.Value) {
	// TODO consider add pushOperand to callStack so stackFrame can be value that is replaced on top.
	vm.callStack.top().push(v)
}
func (vm *VM) pushNewFrame() *stackFrame {
	frame := &stackFrame{env: vm.localEnv().newChild()}
	vm.callStack.push(frame)
	return frame
}
func (vm *VM) popFrame() *stackFrame {
	return vm.callStack.pop()
}
func (vm *VM) fatal(err any) {
	s := structexplorer.NewService("vm", vm)
	for i, each := range vm.callStack {
		s.Explore(fmt.Sprintf("vm.callStack.%d", i), each, structexplorer.Column(0))
	}
	s.Dump("vm-panic.html")
	panic(err)
}

func (vm *VM) eval(e Evaluable) {
	if trace {
		fmt.Println("VM.eval", e)
	}
	e.Eval(vm)
}
