package internal

import "go/ast"

type Field struct {
	*ast.Field
	Names []*Ident
	Type  Expr
}

type FieldList struct {
	*ast.FieldList
	List *Field
}
