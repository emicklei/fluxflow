package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

var _ Stmt = RangeStmt{}

type RangeStmt struct {
	*ast.RangeStmt
	Key, Value Expr // Key, Value may be nil
	X          Expr
	Body       *BlockStmt
}

func (r RangeStmt) String() string {
	return fmt.Sprintf("RangeStmt(%v, %v, %v, %v)", r.Key, r.Value, r.X, r.Body)
}

func (r RangeStmt) stmtStep() Evaluable { return r }

func (r RangeStmt) Eval(vm *VM) {
	rangeable := vm.returnsEval(r.X)
	vm.pushNewFrame()
	for i := 0; i < rangeable.Len(); i++ {
		if r.Key != nil {
			if ca, ok := r.Key.(CanAssign); ok {
				if i == 0 {
					ca.Define(vm, reflect.ValueOf(i))
				} else {
					ca.Assign(vm, reflect.ValueOf(i))
				}
			}
		}
		if r.Value != nil {
			if ca, ok := r.Value.(CanAssign); ok {
				if i == 0 {
					ca.Define(vm, rangeable.Index(i))
				} else {
					ca.Assign(vm, rangeable.Index(i))
				}
			}
		}
		vm.eval(r.Body)
	}
	vm.popFrame()
}

func (r RangeStmt) Flow(g *grapher) (head Step) {
	return head // TODO
}
