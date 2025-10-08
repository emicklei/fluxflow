package internal

import (
	"os"

	"github.com/emicklei/dot"
)

type grapher struct {
	head    Step
	current Step
}

func (g *grapher) next(stmt Evaluable) {
	g.nextStep(newStep(stmt))
}
func (g *grapher) nextStep(next Step) {
	if g.current != nil {
		g.current.SetNext(next)
	} else {
		g.head = next
	}
	g.current = next
}
func (g *grapher) beginIf(cond Evaluable) *conditionalStep {
	c := &conditionalStep{
		step: newStep(cond),
	}
	g.nextStep(c)
	return c
}
func (g *grapher) endIf(beginIf *conditionalStep) {
	nop := newStep(nil) // no-op step
	beginIf.falseStep = nop
	g.nextStep(nop)
}
func (g *grapher) jump(s Step) {
	g.nextStep(s)
}
func (g *grapher) dotify() {
	if g.current == nil {
		return
	}
	d := dot.NewGraph(dot.Directed)
	visited := map[int]dot.Node{}
	g.head.Traverse(d, visited)
	os.WriteFile("graph.dot", []byte(d.String()), 0644)
}
