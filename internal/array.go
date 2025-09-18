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
	rt := builtinTypesMap["int"]
	st := reflect.SliceOf(rt)
	vm.Returns(reflect.MakeSlice(st, 0, 0))
}

func (a ArrayType) String() string {
	return fmt.Sprintf("ArrayType(%v)", a.Elt)
}
