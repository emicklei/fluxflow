package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

var _ Stmt = ReturnStmt{}

type ReturnStmt struct {
	*ast.ReturnStmt
	Results []Expr
}

func (r ReturnStmt) stmtStep() Evaluable { return r }

func (r ReturnStmt) String() string {
	return fmt.Sprintf("return(len=%d)", len(r.Results))
}

func (r ReturnStmt) Eval(vm *VM) {
	// TODO optimize for empty results
	results := make([]reflect.Value, len(r.Results))
	for i, each := range r.Results {
		var val reflect.Value
		if vm.isStepping {
			val = vm.callStack.top().pop()
		} else {
			val = vm.returnsEval(each)
		}
		results[i] = val
	}
	// set return values in the current frame
	frame := vm.callStack.pop()
	frame.returnValues = results
	vm.callStack.push(frame)
}

func (r ReturnStmt) Flow(g *grapher) (head Step) {
	head = g.current
	// reverse order to keep Eval correct
	for i := len(r.Results) - 1; i >= 0; i-- {
		each := r.Results[i]
		if i == 0 {
			head = each.Flow(g)
			continue
		}
		each.Flow(g)
	}
	return
}
