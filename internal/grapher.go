package internal

import (
	"os"

	"github.com/emicklei/dot"
)

type grapher struct {
	head           Step
	current        Step
	conditionStack stack[*conditionalStep]
}

func (g *grapher) next(stmt Evaluable) {
	g.nextStep(newStep(stmt))
}
func (g *grapher) nextStep(next Step) {
	if g.current != nil {
		g.current.SetNext(next)
		return
	}
	g.head = next
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
	g.current.SetNext(s)
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
