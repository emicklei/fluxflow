package internal

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"path"
	"reflect"
	"strconv"
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

func (b *builder) envSet(name string, value reflect.Value) {
	if os.Getenv("STEPS") != "" {
		fmt.Fprintf(os.Stderr, "%s -> %v\n", name, value)
	}
	b.env.set(name, value)
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
		blk := e.(BlockStmt)
		s.Body = &blk
		b.push(s)
	case *ast.UnaryExpr:
		s := &UnaryExpr{UnaryExpr: n}
		b.Visit(n.X)
		e := b.pop()
		s.X = e.(Expr)
		b.push(s)
	case *ast.ValueSpec:
		s := ConstOrVar{spec: n}
		b.Visit(n.Type)
		e := b.pop()
		s.Type = e.(Expr)
		for _, each := range n.Values {
			b.Visit(each)
			e = b.pop()
			s.Values = append(s.Values, e.(Expr))
		}
		b.push(s)
	case *ast.ExprStmt:
		s := ExprStmt{ExprStmt: n}
		b.Visit(n.X)
		e := b.pop()
		s.X = e.(Expr)
		b.push(s)
	case *ast.Ident:
		s := Ident{Ident: n}
		b.push(s)
	case *ast.BlockStmt:
		s := BlockStmt{BlockStmt: n}
		for _, stmt := range n.List {
			b.Visit(stmt)
			e := b.pop()
			s.List = append(s.List, e.(Stmt))
		}
		b.push(s)
	case *ast.AssignStmt:
		s := AssignStmt{AssignStmt: n}
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
		unq, _ := strconv.Unquote(n.Path.Value)
		p := Package{Path: unq}
		if n.Name != nil {
			p.Name = n.Name.Name
		} else {
			// derive name from path
			p.Name = path.Base(unq)
		}
		b.envSet(p.Name, reflect.ValueOf(p))
	case *ast.BasicLit:
		s := BasicLit{BasicLit: n}
		b.push(s)
	case *ast.BinaryExpr:
		s := BinaryExpr{BinaryExpr: n}
		b.Visit(n.X)
		e := b.pop()
		s.X = e.(Expr)
		b.Visit(n.Y)
		e = b.pop()
		s.Y = e.(Expr)
		b.push(s)
	case *ast.CallExpr:
		s := CallExpr{CallExpr: n}
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
		s := SelectorExpr{SelectorExpr: n}
		b.Visit(n.X)
		e := b.pop()
		s.X = e.(Expr)
		b.push(s)
	case *ast.StarExpr:
		s := StarExpr{StarExpr: n}
		b.Visit(n.X)
		e := b.pop()
		s.X = e.(Expr)
		b.push(s)
	case *ast.IfStmt:
		s := IfStmt{IfStmt: n}
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
		blk := e.(BlockStmt)
		s.Body = &blk
		if n.Else != nil {
			b.Visit(n.Else)
			e = b.pop()
			s.Else = e.(Stmt)
		}
		b.push(s)
	case *ast.ReturnStmt:
		s := ReturnStmt{ReturnStmt: n}
		for _, r := range n.Results {
			b.Visit(r)
			e := b.pop()
			s.Results = append(s.Results, e.(Expr))
		}
		b.push(s)
	case *ast.FuncDecl:
		s := FuncDecl{FuncDecl: n}
		if n.Recv != nil {
			b.Visit(n.Recv)
			e := b.pop()
			f := e.(FieldList)
			s.Recv = &f
		}
		b.Visit(n.Name)
		e := b.pop()
		i := e.(Ident)
		s.Name = &i

		b.Visit(n.Type)
		e = b.pop()
		f := e.(FuncType)
		s.Type = &f

		b.Visit(n.Body)
		e = b.pop()
		blk := e.(BlockStmt)
		s.Body = &blk
		b.push(s) // ??
		b.envSet(n.Name.Name, reflect.ValueOf(s))
	case *ast.FuncType:
		s := FuncType{FuncType: n}
		if n.TypeParams != nil {
			b.Visit(n.TypeParams)
			e := b.pop()
			s.TypeParams = e.(*FieldList)
		}
		if n.Params != nil {
			b.Visit(n.Params)
			e := b.pop()
			f := e.(FieldList)
			s.Params = &f
		}
		if n.Results != nil {
			b.Visit(n.Results)
			e := b.pop()
			f := e.(FieldList)
			s.Returns = &f
		}
		b.push(s)
	case *ast.FieldList:
		s := FieldList{FieldList: n}
		for _, field := range n.List {
			b.Visit(field)
			e := b.pop()
			f := e.(Field)
			s.List = append(s.List, &f)
		}
		b.push(s)
	case *ast.Field:
		s := Field{Field: n}
		for _, name := range n.Names {
			b.Visit(name)
			e := b.pop()
			i := e.(Ident)
			s.Names = append(s.Names, &i)
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
		case token.IMPORT:
			for _, each := range n.Specs {
				b.Visit(each)
			}
		case token.TYPE:
			for _, each := range n.Specs {
				b.Visit(each)
			}
		}
	case *ast.DeclStmt:
		s := DeclStmt{DeclStmt: n}
		b.Visit(n.Decl)
		e := b.pop()
		s.Decl = e.(Decl)
		b.push(s)
	case *ast.CompositeLit:
		s := CompositeLit{CompositeLit: n}
		b.Visit(n.Type)
		e := b.pop()
		s.Type = e.(Expr)
		if n.Elts != nil {
			for _, elt := range n.Elts {
				b.Visit(elt)
				e := b.pop()
				s.Elts = append(s.Elts, e.(Expr))
			}
		}
		b.push(s)
	case *ast.ArrayType:
		s := ArrayType{ArrayType: n}
		if n.Len != nil {
			b.Visit(n.Len)
			e := b.pop()
			s.Len = e.(Expr)
		}
		b.Visit(n.Elt)
		e := b.pop()
		s.Elt = e.(Expr)
		b.push(s)
	case *ast.KeyValueExpr:
		s := KeyValueExpr{KeyValueExpr: n}
		b.Visit(n.Key)
		e := b.pop()
		s.Key = e.(Expr)
		b.Visit(n.Value)
		e = b.pop()
		s.Value = e.(Expr)
		b.push(s)
	case *ast.TypeSpec:
		s := TypeSpec{TypeSpec: n}
		if n.Name != nil {
			b.Visit(n.Name)
			e := b.pop().(Ident)
			s.Name = &e
		}
		if n.TypeParams != nil {
			b.Visit(n.TypeParams)
			e := b.pop().(FieldList)
			s.TypeParams = &e
		}
		b.Visit(n.Type)
		e := b.pop().(Expr)
		s.Type = e
		b.push(s)
		if s.Name != nil {
			b.envSet(s.Name.Name, reflect.ValueOf(s))
		} else {
			// what if nil?
		}
	case *ast.StructType:
		s := StructType{StructType: n}
		if n.Fields != nil {
			b.Visit(n.Fields)
			e := b.pop().(FieldList)
			s.Fields = &e
		}
		b.push(s)
	case nil:
		// end of a branch
	default:
		fmt.Fprintf(os.Stderr, "unvisited %T\n", n)
	}
	return b
}
