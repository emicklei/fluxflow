package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

var _ Expr = KeyValueExpr{}

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

func (e KeyValueExpr) Flow(g *grapher) (head Step) {
	head = e.Value.Flow(g)
	e.Key.Flow(g)
	return head
}

type KeyValue struct {
	Key   reflect.Value
	Value reflect.Value
}

func (k KeyValue) String() string {
	return fmt.Sprintf("KeyValue(%v,%v)", k.Key, k.Value)
}
