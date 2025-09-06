package nodes

import "go/ast"

type Call struct {
	node
	*ast.CallExpr
}
