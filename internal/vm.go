package internal

import (
	"bytes"
	"reflect"
)

type stackFrame struct {
	env          ienv
	funcArgs     []reflect.Value
	returnValues []reflect.Value
}

type VM struct {
	operandStack stack[reflect.Value]
	callStack    stack[stackFrame]
	env          ienv
	output       *bytes.Buffer
}

func newVM() *VM {
	return &VM{env: newEnv(), output: new(bytes.Buffer)}
}

func (vm *VM) localEnv() ienv {
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
