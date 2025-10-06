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
		vm.pushOperand(reflect.ValueOf(i))
	case token.STRING:
		unq := strings.Trim(s.Value, "`\"")
		vm.pushOperand(reflect.ValueOf(unq))
	case token.FLOAT:
		f, _ := strconv.ParseFloat(s.Value, 64)
		vm.pushOperand(reflect.ValueOf(f))
	case token.CHAR:
		// a character literal is a rune, which is an alias for int32
		vm.pushOperand(reflect.ValueOf(s.Value))
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
	internalType := vm.returnsEval(s.Type).Interface()
	i, ok := internalType.(CanInstantiate)
	if !ok {
		panic(fmt.Sprintf("expected CanInstantiate:%v (%T)", internalType, internalType))
	}
	instance := i.Instantiate(vm)
	values := make([]reflect.Value, len(s.Elts))
	for i, elt := range s.Elts {
		values[i] = vm.returnsEval(elt)
	}
	result := i.LiteralCompose(instance, values)
	vm.pushOperand(result)
}

func (s CompositeLit) String() string {
	return fmt.Sprintf("CompositeLit(%v,%v)", s.Type, s.Elts)
}

type FuncLit struct {
	*ast.FuncLit
	Type *FuncType
	Body *BlockStmt // TODO not sure what to do when Body and/or Type is nil
}

func (s FuncLit) Eval(vm *VM) {
	vm.pushOperand(reflect.ValueOf(s))
}

func (s FuncLit) String() string {
	return fmt.Sprintf("FuncLit(%v,%v)", s.Type, s.Body)
}
