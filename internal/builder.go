package internal

import "go/ast"

var _ ast.Visitor = (*builder)(nil)

type builder struct {
	root Step
}

// Visit implements the ast.Visitor interface
func (b *builder) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.ValueSpec:
		s := &Var{spec: n}
		b.root.AddStep(s)
	}
	return b
}
