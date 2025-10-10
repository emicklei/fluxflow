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
		g.next(i.Init)
		head = g.current
	}
	end := newStep(nil)
	begin := g.beginIf(i.Cond)
	if head == nil {
		head = begin
	}
	i.Body.Flow(g)
	g.nextStep(end)
	if i.Else != nil {
		begin.elseStep = i.Else.Flow(g)
		g.nextStep(end)
	}
	return head
}
