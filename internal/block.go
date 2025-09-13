package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type BlockStmt struct {
	*ast.BlockStmt
	List []Stmt
}

func (b *BlockStmt) stmtStep() Evaluable { return b }

func (b *BlockStmt) String() string {
	return fmt.Sprintf("BlockStmt(len=%d)", len(b.List))
}

func (b *BlockStmt) Eval(vm *VM) reflect.Value {
	for _, stmt := range b.List {
		stmt.stmtStep().Eval(vm)
	}
	return reflect.Value{}
}
