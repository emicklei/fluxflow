package internal

import "reflect"

type Step interface {
	AddStep(substep Step)
	SetParent(parent Step)
	Eval(env *Env) reflect.Value
}

type step struct {
	parent Step
	steps  []Step
}

func (s step) Eval(env *Env) reflect.Value {
	return reflect.Value{}
}

func (s *step) AddStep(substep Step) {
	substep.SetParent(s)
	s.steps = append(s.steps, substep)
}
func (s *step) SetParent(parent Step) {
	s.parent = parent
}
