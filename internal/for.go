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

	fr := stackFrame{env: vm.localEnv().newChild()}
	vm.callStack.push(fr)

	if f.Init != nil {
		f.Init.stmtStep().Eval(vm)
	}
	for vm.ReturnsEval(f.Cond).Bool() {
		f.Body.stmtStep().Eval(vm)
		f.Post.stmtStep().Eval(vm)
	}

	vm.callStack.pop()
}
