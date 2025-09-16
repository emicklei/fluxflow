package internal

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
	"strings"
)

type BasicLit struct {
	*ast.BasicLit
}

func (s BasicLit) Eval(vm *VM) {
	switch s.Kind {
	case token.INT:
		i, _ := strconv.Atoi(s.Value)
		vm.Returns(reflect.ValueOf(i))
		return
	case token.STRING:
		unq := strings.Trim(s.Value, "`\"")
		vm.Returns(reflect.ValueOf(unq))
		return
	case token.FLOAT:
		f, _ := strconv.ParseFloat(s.Value, 64)
		vm.Returns(reflect.ValueOf(f))
		return
	case token.CHAR:
		// a character literal is a rune, which is an alias for int32
		vm.Returns(reflect.ValueOf(s.Value))
		return
	}
	panic("not implemented: BasicList.Eval:" + s.Kind.String())
}
func (s BasicLit) Loc(f *token.File) string {
	return fmt.Sprintf("%v:BasicLit(%v)", f.Position(s.Pos()), s.Value)
}
func (s BasicLit) String() string {
	return fmt.Sprintf("BasicLit(%v)", s.Value)
}

type CompositeLit struct {
	*ast.CompositeLit
	Elts []Expr
}

func (s CompositeLit) Eval(vm *VM) {}

func (s CompositeLit) String() string {
	return fmt.Sprintf("CompositeLit(%v)", s.Elts)
}
