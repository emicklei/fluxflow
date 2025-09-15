package internal

import (
	"bytes"
	"reflect"
)

type VM struct {
	expressionValue reflect.Value
	callStack       stack
	env             *Env // global or package level?
	output          *bytes.Buffer
}

func newVM() *VM {
	return &VM{env: newEnv(), output: new(bytes.Buffer)}
}

func (vm *VM) localEnv() *Env {
	return vm.env
}

// ReturnsEval evaluates the argument and returns its return value.
func (vm *VM) ReturnsEval(e Evaluable) reflect.Value {
	e.Eval(vm)
	return vm.expressionValue
}
func (vm *VM) Returns(v reflect.Value) {
	vm.expressionValue = v
}
