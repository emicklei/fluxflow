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
		i.Init.Eval(vm)
	}
	rv := vm.ReturnsEval(i.Cond)
	if rv.Bool() {
		i.Body.Eval(vm)
		return
	}
	if i.Else != nil {
		i.Else.stmtStep().Eval(vm)
		return
	}
}
