package internal

import (
	"fmt"
	"go/ast"
)

type ForStmt struct {
	*ast.ForStmt
	Init Stmt
	Cond Expr
	Post Stmt
	Body *BlockStmt
}

func (f ForStmt) stmtStep() Evaluable { return f }

func (f ForStmt) String() string {
	return fmt.Sprintf("ForStmt(%v)", f.Cond)
}
func (f ForStmt) Eval(vm *VM) {
	vm.pushNewFrame()
	if f.Init != nil {
		vm.eval(f.Init.stmtStep())
	}
	for vm.returnsEval(f.Cond).Bool() {
		vm.eval(f.Body.stmtStep())
		vm.eval(f.Post.stmtStep())
	}
	vm.popFrame()
}

func (f ForStmt) Flow(g *grapher) {
	g.next(f.Init.stmtStep())
	begin := g.beginIf(f.Cond)
	f.Body.Flow(g)
	//f.Post.Flow(g)
	g.jump(begin)
	g.endIf(begin)
}
