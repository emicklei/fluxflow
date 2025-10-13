package internal

import (
	"fmt"
	"go/ast"
)

var _ Stmt = IfStmt{}

type IfStmt struct {
	*ast.IfStmt
	Init Expr
	Cond Expr
	Body *BlockStmt
	Else Stmt // else if ...
}

func (i IfStmt) stmtStep() Evaluable { return i }

func (i IfStmt) String() string {
	return fmt.Sprintf("IfStmt(%v, %v, %v, %v)", i.Init, i.Cond, i.Body, i.Else)
}

func (i IfStmt) Eval(vm *VM) {
	if i.Init != nil {
		vm.eval(i.Init)
	}
	rv := vm.returnsEval(i.Cond)
	if rv.Bool() {
		vm.eval(i.Body)
		return
	}
	if i.Else != nil {
		vm.eval(i.Else.stmtStep())
		return
	}
}

func (i IfStmt) Flow(g *grapher) (head Step) {
	if i.Init != nil {
		head = i.Init.Flow(g)
	}
	// this is the flow of cond returing the head
	// this head already has next steps
	cond := i.Cond.Flow(g)
	if i.Init == nil {
		head = cond
	}
	begin := &conditionalStep{
		step: newStep(nil),
	}
	g.nextStep(begin)

	// both true and false branch need a new stack frame
	truePush := &pushStackFrameStep{step: newStep(nil)}
	// both branches will pop and can use the same step
	pop := &popStackFrameStep{step: newStep(nil)}

	g.nextStep(truePush)
	i.Body.Flow(g)
	// after true branch
	g.nextStep(pop)

	// now handle false branch
	if i.Else != nil {
		elsePush := &pushStackFrameStep{step: newStep(nil)}
		begin.elseStep = elsePush
		g.current = elsePush
		i.Else.Flow(g)
		// after false branch
		g.nextStep(pop)
	}
	return head
}
