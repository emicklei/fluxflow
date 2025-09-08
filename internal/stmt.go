package internal

import "go/ast"

type Stmt struct {
	step
	*ast.ExprStmt
}
