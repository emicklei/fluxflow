package internal

import "reflect"

type Step interface {
	Eval(env *Env) reflect.Value
	Next(n Step)
	Prev(p Step)
}

type step struct {
	prev Step
	next Step
}

func (s step) Eval(env *Env) reflect.Value {
	return reflect.Value{}
}

func (s *step) Next(n Step) {
	if s.next == n {
		return
	}
	s.next = n
	n.Prev(s)
}
func (s *step) Prev(p Step) {
	if s.prev == p {
		return
	}
	s.prev = p
	p.Next(s)
}
