package internal

import (
	"fmt"
)

var _ Step = (*step)(nil)

var idgen int = 0

type step struct {
	id   int
	prev Step
	next Step
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

func (s *step) ID() int {
	return s.id
}

func (s *step) String() string {
	if s == nil {
		return "nil"
	}
	return fmt.Sprintf("step(%v)", s.Evaluable)
}

func (s *step) Next() Step {
	return s.next
}
func (s *step) Prev() Step {
	return s.prev
}

func (s *step) SetNext(n Step) {
	if n == s {
		panic("step cannot point to itself")
	}
	if s.next == n {
		return
	}
	s.next = n
	n.SetPrev(s)
}
func (s *step) SetPrev(p Step) {
	if p == s {
		panic("step cannot point to itself")
	}
	if s.prev == p {
		return
	}
	s.prev = p
	p.SetNext(s)
}
