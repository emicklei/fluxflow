package internal

import (
	"fmt"
	"go/ast"
	"log"
)

var _ ast.Visitor = (*builder)(nil)

type builder struct {
	stack []any
}

func (b *builder) push(s any) {
	fmt.Printf("%T\n", s)
	b.stack = append(b.stack, s)
}
func (b *builder) pop() any {
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
			e := b.pop()
			s.Lhs = append(s.Lhs, e.(Expr))
		}
		for _, r := range n.Rhs {
			b.Visit(r)
			e := b.pop()
			s.Rhs = append(s.Rhs, e.(Expr))
		}
		b.push(s)
	case *ast.ImportSpec:
	case *ast.BasicLit:
		s := &BasicLit{BasicLit: n}
		b.push(s)
	case *ast.BinaryExpr:
		s := &BinaryExpr{BinaryExpr: n}
		b.Visit(n.X)
		e := b.pop()
		s.X = e.(Expr)
		b.Visit(n.Y)
		e = b.pop()
		s.Y = e.(Expr)
		b.push(s)
	default:
		log.Println("unvisited", fmt.Sprintf("%T", n))
	}
	return b
}
