package internal

type grapher struct {
	current        *step
	conditionStack stack[*conditionalStep]
}

func (g *grapher) next(stmt Evaluable) {
	next := &step{Evaluable: stmt}
	if g.current != nil {
		g.current.Next(next)
	}
	g.current = next
}
func (g *grapher) beginIf(cond Evaluable) *step {
	c := &conditionalStep{
		step: &step{Evaluable: cond},
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
