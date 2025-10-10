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
		results[i] = vm.returnsEval(each)
	}
	frame := vm.callStack.pop()
	frame.returnValues = results
	vm.callStack.push(frame)
}

func (r ReturnStmt) Flow(g *grapher) (head Step) {
	head = g.current
	for i, each := range r.Results {
		if i == 0 {
			head = each.Flow(g)
			continue
		}
		each.Flow(g)
	}
	return
}
