package internal

import (
	"fmt"
	"go/ast"
)

var _ ast.Visitor = (*builder)(nil)

type builder struct {
	stack []*step
}

func (b *builder) push(s Evaluable) {
	if str, ok := s.(fmt.Stringer); ok {
		fmt.Printf("%v\n", str.String())
	} else {
		fmt.Printf("%T\n", s)
	}
	step := new(step)
	step.Evaluable = s
	if len(b.stack) > 0 {
		top := b.stack[len(b.stack)-1]
		step.Prev(top)
	}
	b.stack = append(b.stack, step)
}

func (b *builder) pop() Evaluable {
	top := b.stack[len(b.stack)-1]
	b.stack = b.stack[0 : len(b.stack)-1]
	return top.Evaluable
}

func (b *builder) last() *step {
	if len(b.stack) == 0 {
		return nil
	}
	return b.stack[len(b.stack)-1]
}
func (b *builder) first() *step {
	if len(b.stack) == 0 {
		return nil
	}
	return b.stack[0]
}

// Visit implements the ast.Visitor interface
func (b *builder) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.UnaryExpr:
		s := &UnaryExpr{UnaryExpr: n}
		b.Visit(n.X)
		e := b.pop()
		s.X = e.(Expr)
		b.push(s)
	case *ast.ValueSpec:
		s := &Var{spec: n}
		b.push(s)
	case *ast.ExprStmt:
		s := &ExprStmt{ExprStmt: n}
		b.Visit(n.X)
		e := b.pop()
		s.X = e.(Expr)
		b.push(s)
	case *ast.Ident:
		s := &Ident{Ident: n}
		b.push(s)
	case *ast.BlockStmt:
		s := &BlockStmt{BlockStmt: n}
		for _, stmt := range n.List {
			b.Visit(stmt)
			e := b.pop()
			s.List = append(s.List, e.(Stmt))
		}
		b.push(s)
	case *ast.AssignStmt:
		s := &AssignStmt{AssignStmt: n}
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
	case *ast.CallExpr:
		s := &CallExpr{CallExpr: n}
		b.Visit(n.Fun)
		e := b.pop()
		s.Fun = e.(Expr)
		for _, arg := range n.Args {
			b.Visit(arg)
			e := b.pop()
			s.Args = append(s.Args, e.(Expr))
		}
		b.push(s)
	case *ast.SelectorExpr:
		s := &SelectorExpr{SelectorExpr: n}
		b.Visit(n.X)
		e := b.pop()
		s.X = e.(Expr)
		b.push(s)
	case *ast.StarExpr:
		s := &StarExpr{StarExpr: n}
		b.Visit(n.X)
		e := b.pop()
		s.X = e.(Expr)
		b.push(s)
	case *ast.IfStmt:
		s := &IfStmt{IfStmt: n}
		if n.Init != nil {
			b.Visit(n.Init)
			e := b.pop()
			s.Init = e.(Expr)
		}
		b.Visit(n.Cond)
		e := b.pop()
		s.Cond = e.(Expr)
		b.Visit(n.Body)
		e = b.pop()
		s.Body = e.(*BlockStmt)
		b.push(s)
	case *ast.ReturnStmt:
		s := &ReturnStmt{ReturnStmt: n}
		for _, r := range n.Results {
			b.Visit(r)
			e := b.pop()
			s.Results = append(s.Results, e.(Expr))
		}
		b.push(s)
	case *ast.FuncDecl:
		s := &FuncDecl{FuncDecl: n}
		b.Visit(n.Name)
		e := b.pop()
		s.Name = e.(*Ident)
		b.Visit(n.Body)
		e = b.pop()
		s.Body = e.(*BlockStmt)
		b.push(s)
	case *ast.GenDecl:
		// IMPORT, CONST, TYPE, or VAR
	default:
		fmt.Println("unvisited", fmt.Sprintf("%T", n))
	}
	return b
}
