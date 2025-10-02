package internal

import (
	"bytes"
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
	e.Eval(vm)
	return vm.operandStack.pop()
}
func (vm *VM) Returns(v reflect.Value) {
	vm.operandStack.push(v)
}
