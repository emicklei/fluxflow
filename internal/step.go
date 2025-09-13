package internal

import (
	"fmt"
)

type step struct {
	prev *step
	next *step
	Evaluable
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
