package internal

type VM struct {
	callStack stack
	env       *Env
}

func newVM() *VM {
	return &VM{env: newEnv()}
}
