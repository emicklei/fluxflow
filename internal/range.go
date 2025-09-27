package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

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
	rangeable := vm.ReturnsEval(r.X)
	frame := stackFrame{env: vm.env.subEnv()}
	vm.callStack.push(frame)
	for i := 0; i < rangeable.Len(); i++ {
		if r.Key != nil {
			if ca, ok := r.Key.(CanAssign); ok {
				if i == 0 {
					ca.Define(vm.localEnv(), reflect.ValueOf(i))
				} else {
					ca.Assign(vm.localEnv(), reflect.ValueOf(i))
				}
			}
		}
		if r.Value != nil {
			if ca, ok := r.Value.(CanAssign); ok {
				if i == 0 {
					ca.Define(vm.localEnv(), rangeable.Index(i))
				} else {
					ca.Assign(vm.localEnv(), rangeable.Index(i))
				}
			}
		}
		r.Body.Eval(vm)
	}
	vm.callStack.pop()
}
