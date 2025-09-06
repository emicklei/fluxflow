package nodes

import "go/ast"

type Var struct {
	node
	*ast.ValueSpec
}
