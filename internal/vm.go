package internal

type VM struct {
	callStack stack
	env       *Env
}

func newVM() *VM {
	return &VM{env: newEnv()}
}

func (vm *VM) localEnv() *Env {
	return vm.env
}
