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
	rt := builtinTypesMap["int"] // TODO support other types than int
	st := reflect.SliceOf(rt)
	vm.Returns(reflect.MakeSlice(st, 0, 0))
}

func (a ArrayType) String() string {
	return fmt.Sprintf("ArrayType(%v)", a.Elt)
}

func (a ArrayType) LiteralCompose(composite reflect.Value, values ...reflect.Value) reflect.Value {
	for _, v := range values {
		composite = reflect.Append(composite, v)
	}
	return composite
}
