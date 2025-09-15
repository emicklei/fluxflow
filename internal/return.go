package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type ReturnStmt struct {
	*ast.ReturnStmt
	Results []Expr
}

func (r *ReturnStmt) stmtStep() Evaluable { return r }

func (r *ReturnStmt) String() string {
	return fmt.Sprintf("return(len=%d)", len(r.Results))
}

func (r *ReturnStmt) Eval(vm *VM) {
	results := make([]reflect.Value, len(r.Results))
	for i, each := range r.Results {
		results[i] = vm.ReturnsEval(each)
	}
	frame := vm.callStack.pop()
	frame.returnValues = results
	vm.callStack.push(frame)
}
