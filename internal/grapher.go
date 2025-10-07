package internal

import (
	"os"
	"strconv"

	"github.com/emicklei/dot"
)

type grapher struct {
	current        *step
	conditionStack stack[*conditionalStep]
}

func (g *grapher) next(stmt Evaluable) {
	next := newStep(stmt)
	if g.current != nil {
		g.current.Next(next)
	}
	g.current = next
}
func (g *grapher) beginIf(cond Evaluable) *step {
	c := &conditionalStep{
		step: newStep(cond),
	}
	g.next(c)
	g.conditionStack.push(c)
	return c.step // nicht gut
}
func (g *grapher) endIf() {
	nop := &step{Evaluable: nil} // no-op step
	c := g.conditionStack.pop()
	c.falseStep = nop
	g.next(nop)
}
func (g *grapher) jump(s *step) {
	g.current.Next(s)
}
func (g *grapher) dotify() {
	if g.current == nil {
		return
	}
	d := dot.NewGraph(dot.Directed)
	here := g.current.head()
	hereNode := d.Node(strconv.Itoa(here.id)).Label(here.String())
	for ; here != nil; here = here.next {
		nextNode := d.Node(strconv.Itoa(here.id))
		if nextNode.HasDefaultLabel() {
			nextNode.Label(here.String())
		}
		hereNode.Edge(nextNode, "")
		hereNode = nextNode
	}
	os.WriteFile("graph.dot", []byte(d.String()), 0644)
}
