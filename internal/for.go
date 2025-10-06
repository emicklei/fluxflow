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
	for vm.ReturnsEval(f.Cond).Bool() {
		vm.eval(f.Body.stmtStep())
		vm.eval(f.Post.stmtStep())
	}
	vm.popFrame()
}
