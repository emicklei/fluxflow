package internal

import (
	"fmt"
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
func (s BasicLit) Loc(f *token.File) string {
	return fmt.Sprintf("%v:BasicLit(%v)", f.Position(s.Pos()), s.Value)
}
func (s BasicLit) String() string {
	return fmt.Sprintf("BasicLit(%v)", s.Value)
}
