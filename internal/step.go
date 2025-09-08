package internal

type Step interface {
	AddStep(substep Step)
	SetParent(parent Step)
}

type step struct {
	parent Step
	steps  []Step
}

func (s *step) AddStep(substep Step) {
	substep.SetParent(s)
	s.steps = append(s.steps, substep)
}
func (s *step) SetParent(parent Step) {
	s.parent = parent
}
