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
	begin := g.beginIf(i.Cond)
	// if no init then head is begin
	if head == nil {
		head = begin
	}
	// both true and false branch need a new stack frame
	push := &pushStackFrameStep{step: newStep(nil)}
	// both branches will pop and can use the same step
	pop := &popStackFrameStep{step: newStep(nil)}

	g.nextStep(push)
	i.Body.Flow(g)
	// after true branch
	g.nextStep(pop)

	// now handle false branch
	if i.Else != nil {
		elsePush := &pushStackFrameStep{step: newStep(nil)}
		// branching to else
		g.current = elsePush
		begin.elseStep = elsePush
		i.Else.Flow(g)
		// after false branch
		g.nextStep(pop)
	}
	return head
}
