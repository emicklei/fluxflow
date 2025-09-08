package internal

import "go/ast"

// For both variable and constant for now
type Var struct {
	step
	spec *ast.ValueSpec
}
