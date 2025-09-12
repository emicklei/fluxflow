package internal

import (
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
)

type BasicLit struct {
	operatorUnimplemented
	Step
	*ast.BasicLit
}

func (s BasicLit) Eval(env *Env) reflect.Value {
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
