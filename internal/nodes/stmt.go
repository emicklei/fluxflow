package nodes

import "go/ast"

type Stmt struct {
	node
	*ast.ExprStmt
}
