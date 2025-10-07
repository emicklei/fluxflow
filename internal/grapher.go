package internal

import (
	"os"
	"strconv"

	"github.com/emicklei/dot"
)

type grapher struct {
	current        Step
	conditionStack stack[*conditionalStep]
}

func (g *grapher) next(stmt Evaluable) {
	g.nextStep(newStep(stmt))
}
func (g *grapher) nextStep(next Step) {
	if g.current != nil {
		g.current.SetNext(next)
	}
	g.current = next
}
func (g *grapher) beginIf(cond Evaluable) Step {
	c := &conditionalStep{
		step: newStep(cond),
	}
	g.nextStep(c)
	g.conditionStack.push(c)
	return c
}
func (g *grapher) endIf() {
	nop := newStep(nil) // no-op step
	c := g.conditionStack.pop()
	c.falseStep = nop
	g.nextStep(nop)
}
func (g *grapher) jump(s Step) {
	nop := newStep(nil) // no-op step
	g.current.SetNext(nop)
	nop.SetNext(s)
}
func (g *grapher) dotify() {
	if g.current == nil {
		return
	}
	d := dot.NewGraph(dot.Directed)
	// find head
	here := g.current
	for here.Prev() != nil {
		here = here.Prev()
	}
	hereNode := d.Node(strconv.Itoa(here.ID())).Label(here.String())
	for ; here != nil; here = here.Next() {
		nextNode := d.Node(strconv.Itoa(here.ID()))
		if nextNode.HasDefaultLabel() {
			nextNode.Label(here.String())
		}
		hereNode.Edge(nextNode, "")
		hereNode = nextNode
	}
	os.WriteFile("graph.dot", []byte(d.String()), 0644)
}
