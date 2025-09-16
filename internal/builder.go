package internal

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"reflect"
)

var _ ast.Visitor = (*builder)(nil)

type builder struct {
	stack []*step
	env   *Env
}

func (b *builder) push(s Evaluable) {
	if os.Getenv("STEPS") != "" {
		if str, ok := s.(fmt.Stringer); ok {
			fmt.Fprintf(os.Stderr, "%v\n", str.String())
		} else {
			fmt.Fprintf(os.Stderr, "%T\n", s)
		}
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
	if len(b.stack) == 0 {
		panic("builder.stack is empty")
	}
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
	case *ast.IncDecStmt:
		s := &IncDecStmt{IncDecStmt: n}
		b.Visit(n.X)
		e := b.pop()
		s.X = e.(Expr)
		b.push(s)
	case *ast.ForStmt:
		s := &ForStmt{ForStmt: n}
		if n.Init != nil {
			b.Visit(n.Init)
			e := b.pop()
			s.Init = e.(Stmt)
		}
		if n.Cond != nil {
			b.Visit(n.Cond)
			e := b.pop()
			s.Cond = e.(Expr)
		}
		if n.Post != nil {
			b.Visit(n.Post)
			e := b.pop()
			s.Post = e.(Stmt)
		}
		b.Visit(n.Body)
		e := b.pop()
		s.Body = e.(*BlockStmt)
		b.push(s)
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
		if n.Else != nil {
			b.Visit(n.Else)
			e = b.pop()
			s.Else = e.(Stmt)
		}
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
		if n.Recv != nil {
			b.Visit(n.Recv)
			e := b.pop()
			s.Recv = e.(*FieldList)
		}
		b.Visit(n.Name)
		e := b.pop()
		s.Name = e.(*Ident)

		b.Visit(n.Type)
		e = b.pop()
		s.Type = e.(*FuncType)

		b.Visit(n.Body)
		e = b.pop()
		s.Body = e.(*BlockStmt)
		b.push(s)
		// put in scope TODO
		b.env.set(n.Name.Name, reflect.ValueOf(s))
	case *ast.FuncType:
		s := &FuncType{FuncType: n}
		if n.TypeParams != nil {
			b.Visit(n.TypeParams)
			e := b.pop()
			s.TypeParams = e.(*FieldList)
		}
		if n.Params != nil {
			b.Visit(n.Params)
			e := b.pop()
			s.Params = e.(*FieldList)
		}
		if n.Results != nil {
			b.Visit(n.Results)
			e := b.pop()
			s.Returns = e.(*FieldList)
		}
		b.push(s)
	case *ast.FieldList:
		s := &FieldList{FieldList: n}
		for _, field := range n.List {
			b.Visit(field)
			e := b.pop()
			s.List = append(s.List, e.(*Field))
		}
		b.push(s)
	case *ast.Field:
		s := &Field{Field: n}
		for _, name := range n.Names {
			b.Visit(name)
			e := b.pop()
			s.Names = append(s.Names, e.(*Ident))
		}
		b.Visit(n.Type)
		e := b.pop()
		s.Type = e.(Expr)
		// TODO tag, comment
		b.push(s)
	case *ast.GenDecl:
		// IMPORT, CONST, TYPE, or VAR
		switch n.Tok {
		case token.VAR:
			for _, each := range n.Specs {
				b.Visit(each)
			}
		}
	case *ast.DeclStmt:
		s := &DeclStmt{DeclStmt: n}
		b.Visit(n.Decl)
		e := b.pop()
		s.Decl = e.(Decl)
		b.push(s)
	case *ast.CompositeLit:
		s := &CompositeLit{CompositeLit: n}
		if n.Elts != nil {
			for _, elt := range n.Elts {
				b.Visit(elt)
				e := b.pop()
				s.Elts = append(s.Elts, e.(Expr))
			}
		}
		b.push(s)
	case nil:
		// end of a branch
	default:
		fmt.Fprintf(os.Stderr, "unvisited %T\n", n)
	}
	return b
}
