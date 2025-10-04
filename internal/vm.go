package internal

import (
	"bytes"
	"fmt"
	"reflect"
)

type stackFrame struct {
	env          Env
	operandStack stack[reflect.Value]
	returnValues []reflect.Value
}

func (f *stackFrame) push(v reflect.Value) {
	f.operandStack.push(v)
}
func (f *stackFrame) pop() reflect.Value {
	return f.operandStack.pop()
}

type VM struct {
	callStack stack[*stackFrame]
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
	e.Eval(vm)
	return vm.callStack.top().pop()
}

// Returns pushes a value onto the operand stack as the result of an evaluation.
func (vm *VM) Returns(v reflect.Value) {
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
