package internal

import (
	"fmt"
	"go/ast"
)

type BlockStmt struct {
	*ast.BlockStmt
	List []Stmt
}

func (b BlockStmt) stmtStep() Evaluable { return b }

func (b BlockStmt) String() string {
	return fmt.Sprintf("BlockStmt(len=%d)", len(b.List))
}

func (b BlockStmt) Eval(vm *VM) {
	for _, stmt := range b.List {
		vm.eval(stmt.stmtStep())
	}
}

func (b BlockStmt) Flow(g *grapher) {
	// for _, stmt := range b.List {
	// 	stmt.stmtStep().Flow(g)
	// }
	g.next(b)
}
