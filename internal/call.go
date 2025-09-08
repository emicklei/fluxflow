package internal

import "go/ast"

type Call struct {
	node
	*ast.CallExpr
}
