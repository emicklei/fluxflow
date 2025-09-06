package nodes

import "go/ast"

type Func struct {
	node
	*ast.FuncDecl
}
