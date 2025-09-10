package internal

import (
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
)

type Literal struct {
	Step
	*ast.BasicLit
}

func (s Literal) Eval() reflect.Value {
	switch s.Kind {
	case token.INT:
		i, _ := strconv.Atoi(s.Value)
		return reflect.ValueOf(i)
	case token.STRING:
		return reflect.ValueOf(s.Value)
	}
	panic("not implemented")
}
func (s Literal) Operand() Operand {
	switch s.Kind {
	case token.INT:
		return INT{value: s.Value}
	}
	panic("not implemented")
}
