package internal

import (
	"go/ast"
)

type Call struct {
	Step
	*ast.CallExpr
}
