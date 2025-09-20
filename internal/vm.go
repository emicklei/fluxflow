package internal

import (
	"bytes"
	"reflect"
)

type VM struct {
	lastOperand reflect.Value
	callStack   stack
	env         *Env // global or package level?
	output      *bytes.Buffer
}

func newVM() *VM {
	return &VM{env: newEnv(), output: new(bytes.Buffer)}
}

func (vm *VM) localEnv() *Env {
	return vm.callStack.top().env
}

// ReturnsEval evaluates the argument and returns its return value.
func (vm *VM) ReturnsEval(e Evaluable) reflect.Value {
	e.Eval(vm)
	return vm.lastOperand
}
func (vm *VM) Returns(v reflect.Value) {
	vm.lastOperand = v
}
