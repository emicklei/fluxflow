package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type KeyValueExpr struct {
	*ast.KeyValueExpr
	Key   Expr
	Value Expr
}

func (e KeyValueExpr) String() string {
	return fmt.Sprintf("KeyValueExpr(%v,%v)", e.Key, e.Value)
}

func (e KeyValueExpr) Eval(vm *VM) {
	key := vm.ReturnsEval(e.Key)
	value := vm.ReturnsEval(e.Value)
	vm.Returns(reflect.ValueOf(KeyValue{Key: key, Value: value}))
}

type KeyValue struct {
	Key   reflect.Value
	Value reflect.Value
}
