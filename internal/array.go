package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type ArrayType struct {
	*ast.ArrayType
	Elt Expr
}

func (a ArrayType) Eval(vm *VM) {
	// typeName := a.Elt.(Ident).Name
	// TODO
	var i []reflect.Value
	rt := reflect.TypeOf(i)
	vm.Returns(reflect.MakeSlice(rt, 0, 0))
}

func (a ArrayType) String() string {
	return fmt.Sprintf("ArrayType(%v)", a.Elt)
}
