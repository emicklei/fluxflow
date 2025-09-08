package internal

import "go/ast"

type Assign struct {
	step
	*ast.AssignStmt
}
