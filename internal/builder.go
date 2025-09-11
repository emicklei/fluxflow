package internal

import (
	"fmt"
	"go/ast"
	"log"
)

var _ ast.Visitor = (*builder)(nil)

type builder struct {
	stack []Step
	vm    *VM
}

func (b *builder) push(s Step) {
	b.stack = append(b.stack, s)
}
func (b *builder) pop() Step {
	top := b.stack[len(b.stack)-1]
	b.stack = b.stack[0 : len(b.stack)-1]
	return top
}

// Visit implements the ast.Visitor interface
func (b *builder) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.ValueSpec:
		s := &Var{spec: n}
		b.push(s)
	case *ast.ExprStmt:
		s := &Stmt{ExprStmt: n}
		b.push(s)
	case *ast.Ident:
		s := &Ident{Ident: n}
		b.push(s)
	case *ast.BlockStmt:
		for _, s := range n.List {
			b.Visit(s)
		}
	case *ast.AssignStmt:
		s := &Assign{AssignStmt: n}
		for _, l := range n.Lhs {
			b.Visit(l)
		}
		b.push(s)
	case *ast.ImportSpec:
	default:
		log.Println("unvisited", fmt.Sprintf("%T", n))
	}
	return b
}
