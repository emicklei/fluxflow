package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type ArrayType struct {
	*ast.ArrayType
	Len Expr
	Elt Expr
}

func (a ArrayType) Eval(vm *VM) {
	rt := builtinTypesMap["int"] // TODO support other types than int
	if a.ArrayType.Elt == nil {
		st := reflect.SliceOf(rt)
		vm.Returns(reflect.MakeSlice(st, 0, 4))
	} else {
		size := vm.ReturnsEval(a.Len)
		st := reflect.ArrayOf(int(size.Int()), rt)
		pArray := reflect.New(st)
		vm.Returns(pArray.Elem())
	}
}

func (a ArrayType) String() string {
	return fmt.Sprintf("ArrayType(%v,slice=%v)", a.Elt, a.ArrayType.Len == nil)
}

func (a ArrayType) LiteralCompose(composite reflect.Value, values ...reflect.Value) reflect.Value {
	if a.ArrayType.Len == nil { // slice
		for _, v := range values {
			composite = reflect.Append(composite, v)
		}
		return composite
	}
	for i, v := range values {
		composite.Index(i).Set(v)
	}
	return composite
}
