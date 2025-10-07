package internal

import (
	"fmt"
)

var idgen int = 0

type step struct {
	id   int
	prev *step
	next *step
	Evaluable
}

func newStep(e Evaluable) *step {
	idgen++
	return &step{id: idgen, Evaluable: e}
}

type conditionalStep struct {
	*step
	falseStep *step
}

func (s *step) String() string {
	if s == nil {
		return "nil"
	}
	return fmt.Sprintf("step(%v)", s.Evaluable)
}

func (s *step) Next(n *step) {
	if s.next == n {
		return
	}
	s.next = n
	n.Prev(s)
}
func (s *step) Prev(p *step) {
	if s.prev == p {
		return
	}
	s.prev = p
	p.Next(s)
}
func (s *step) head() *step {
	here := s
	for here.prev != nil {
		here = here.prev
	}
	return here
}
