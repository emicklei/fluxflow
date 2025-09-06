package nodes

import "go/ast"

type Assign struct {
	node
	*ast.AssignStmt
}

func (a *Assign) Eval(ec EvalContext) {

}
