package internal

import "bytes"

type VM struct {
	callStack stack
	env       *Env // global or package level?
	output    *bytes.Buffer
}

func newVM() *VM {
	return &VM{env: newEnv(), output: new(bytes.Buffer)}
}

func (vm *VM) localEnv() *Env {
	return vm.env
}
