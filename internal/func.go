package internal

import "go/ast"

type Func struct {
	Step
	*ast.FuncDecl
}
