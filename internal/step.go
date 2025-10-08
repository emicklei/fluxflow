package internal

import (
	"fmt"
	"strconv"

	"github.com/emicklei/dot"
)

var _ Step = (*step)(nil)

var idgen int = 0

type step struct {
	id   int
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

func (c *conditionalStep) Traverse(g *dot.Graph, visited map[int]dot.Node) dot.Node {
	me := c.step.Traverse(g, visited)
	if c.falseStep != nil {
		falseN := c.falseStep.Traverse(g, visited)
		me.Edge(falseN, "false")
	}
	return me
}

func (s *step) Traverse(g *dot.Graph, visited map[int]dot.Node) dot.Node {
	if n, ok := visited[s.id]; ok {
		return n
	}
	n := g.Node(strconv.Itoa(s.ID())).Label(fmt.Sprintf("%d: %s", s.ID(), s.String()))
	visited[s.id] = n
	if s.next != nil {
		nextN := s.next.Traverse(g, visited)
		n.Edge(nextN, "next")
	}
	return n
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

func (s *step) SetNext(n Step) {
	if s.next == n {
		return
	}
	s.next = n
}
