package internal

import (
	"fmt"
	"os"

	"github.com/emicklei/dot"
)

// grapher helps building a control flow graph by keeping track of the current step.
type grapher struct {
	head    Step   // the entry point to the flow graph
	current Step   // the current step to attach the next step to
	dotFile string // for overriding the default graph.dot file (only used when calling dotify)
}

// next adds a new step after the current one and makes it the current step.
func (g *grapher) next(e Evaluable) {
	g.nextStep(newStep(e))
}

// nextStep adds the given step after the current one and makes it the current step.
func (g *grapher) nextStep(next Step) {
	if g.current != nil {
		if g.current.Next() != nil {
			panic(fmt.Sprintf("current %s already has a next %s, failing %s", g.current, g.current.Next(), next))
		}
		g.current.SetNext(next)
	} else {
		g.head = next
	}
	g.current = next
}

// beginIf creates a conditional step with the given condition
// and makes it the current step. It returns the created conditional step to set the else branch later.
func (g *grapher) beginIf(cond Evaluable) *conditionalStep {
	c := &conditionalStep{
		step: newStep(cond),
	}
	g.nextStep(c)
	return c
}

// TODO not sure if needed, replace with pop?
func (g *grapher) endIf(beginIf *conditionalStep) {
	nop := newStep(nil) // no-op step
	beginIf.elseStep = nop
	g.current = nop
}

func (g *grapher) dotFilename() string {
	if g.dotFile != "" {
		return g.dotFile
	}
	return "graph.dot"
}

// dotify writes the current graph to graph.dot
func (g *grapher) dotify() {
	if g.current == nil {
		return
	}
	d := dot.NewGraph(dot.Directed)
	visited := map[int]dot.Node{}
	g.head.Traverse(d, visited)
	os.WriteFile(g.dotFilename(), []byte(d.String()), 0644)
}
