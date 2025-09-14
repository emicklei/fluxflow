package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type ForStmt struct {
	*ast.ForStmt
	Init Stmt
	Cond Expr
	Post Stmt
	Body *BlockStmt
}

func (f *ForStmt) stmtStep() Evaluable { return f }

func (f *ForStmt) String() string {
	return fmt.Sprintf("ForStmt(%v)", f.Cond)
}
func (f *ForStmt) Eval(vm *VM) reflect.Value {
	if f.Init != nil {
		f.Init.stmtStep().Eval(vm)
	}
	return reflect.Value{}
}
