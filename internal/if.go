package internal

import (
	"fmt"
	"go/ast"
)

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
