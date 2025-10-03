package internal

import (
	"bytes"
	"fmt"
	"reflect"
)

type stackFrame struct {
	env          Env
	returnValues []reflect.Value
}

type VM struct {
	operandStack stack[reflect.Value]
	callStack    stack[stackFrame]
	output       *bytes.Buffer
}

func newVM(env Env) *VM {
	vm := &VM{output: new(bytes.Buffer)}
	frame := stackFrame{env: env}
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
	return vm.operandStack.pop()
}

// Returns pushes a value onto the operand stack as the result of an evaluation.
func (vm *VM) Returns(v reflect.Value) {
	vm.operandStack.push(v)
}
func (vm *VM) pushNewFrame() stackFrame {
	frame := stackFrame{env: vm.localEnv().newChild()}
	vm.callStack.push(frame)
	return frame
}
func (vm *VM) popFrame() stackFrame {
	return vm.callStack.pop()
}
