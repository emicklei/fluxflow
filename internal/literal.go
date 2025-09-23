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
	case token.STRING:
		unq := strings.Trim(s.Value, "`\"")
		vm.Returns(reflect.ValueOf(unq))
	case token.FLOAT:
		f, _ := strconv.ParseFloat(s.Value, 64)
		vm.Returns(reflect.ValueOf(f))
	case token.CHAR:
		// a character literal is a rune, which is an alias for int32
		vm.Returns(reflect.ValueOf(s.Value))
	default:
		panic("not implemented: BasicList.Eval:" + s.Kind.String())
	}
}
func (s BasicLit) Loc(f *token.File) string {
	return fmt.Sprintf("%v:BasicLit(%v)", f.Position(s.Pos()), s.Value)
}
func (s BasicLit) String() string {
	return fmt.Sprintf("BasicLit(%v)", s.Value)
}

type CompositeLit struct {
	*ast.CompositeLit
	Type Expr
	Elts []Expr
}

func (s CompositeLit) Eval(vm *VM) {
	internalType := vm.ReturnsEval(s.Type).Interface()
	i, ok := internalType.(CanInstantiate)
	if !ok {
		panic(fmt.Sprintf("expected CanInstantiate:%v (%T)", internalType, internalType))
	}
	composite := i.Instantiate(vm)
	composer, ok := internalType.(CanCompose)
	if !ok {
		panic(fmt.Sprintf("expected CanCompose:%v (%T)", s.Type, s.Type))
	}
	values := make([]reflect.Value, len(s.Elts))
	for i, elt := range s.Elts {
		values[i] = vm.ReturnsEval(elt)
	}
	result := composer.LiteralCompose(composite, values)
	vm.Returns(result)
}

func (s CompositeLit) String() string {
	return fmt.Sprintf("CompositeLit(%v,%v)", s.Type, s.Elts)
}
