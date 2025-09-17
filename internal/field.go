package internal

import (
	"fmt"
	"go/ast"
)

type Field struct {
	*ast.Field
	Names []*Ident
	Type  Expr
}

func (l Field) String() string {
	return fmt.Sprintf("Field(%v,%v)", l.Names, l.Type)
}
func (l Field) Eval(vm *VM) {}

type FieldList struct {
	*ast.FieldList
	List []*Field
}

func (l FieldList) String() string {
	return fmt.Sprintf("FieldList(%v)", l.List)
}
func (l FieldList) Eval(vm *VM) {}
