package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type IfStmt struct {
	*ast.IfStmt
	Init Expr
	Cond Expr
	Body *BlockStmt
	Else *IfStmt // else if ...
}

func (i *IfStmt) stmtStep() Evaluable { return i }

func (i *IfStmt) String() string {
	return fmt.Sprintf("IfStmt(%v, %v, %v, %v)", i.Init, i.Cond, i.Body, i.Else)
}

func (i *IfStmt) Eval(vm *VM) reflect.Value {
	rv := i.Cond.Eval(vm)
	if rv.Bool() {
		return i.Body.Eval(vm)
	}
	if i.Else != nil {
		return i.Else.Eval(vm)
	}
	return reflect.Value{}
}
