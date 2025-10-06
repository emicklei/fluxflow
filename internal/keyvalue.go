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
	id, ok := e.Key.(Ident)
	if !ok {
		panic("unhandled key type")
	}
	key := reflect.ValueOf(id.Name)
	value := vm.returnsEval(e.Value)
	vm.pushOperand(reflect.ValueOf(KeyValue{Key: key, Value: value}))
}

type KeyValue struct {
	Key   reflect.Value
	Value reflect.Value
}

func (k KeyValue) String() string {
	return fmt.Sprintf("KeyValue(%v,%v)", k.Key, k.Value)
}
