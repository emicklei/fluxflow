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

func (s Literal) Eval(env *Env) reflect.Value {
	switch s.Kind {
	case token.INT:
		i, _ := strconv.Atoi(s.Value)
		return reflect.ValueOf(i)
	case token.STRING:
		return reflect.ValueOf(s.Value)
	case token.FLOAT:
		f, _ := strconv.ParseFloat(s.Value, 64)
		return reflect.ValueOf(f)
	case token.CHAR:
		// a character literal is a rune, which is an alias for int32
		return reflect.ValueOf(s.Value)
	}
	panic("not implemented")
}
func (s Literal) Operand() Expr {
	switch s.Kind {
	case token.INT:
		return INT{value: s.Value}
	case token.STRING:
		return STRING{value: s.Value}
	case token.FLOAT:
		return FLOAT{value: s.Value}
	}
	panic("not implemented")
}
