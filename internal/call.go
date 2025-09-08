package internal

import (
	"go/ast"
)

type Call struct {
	Step
	*ast.CallExpr
}

func (s *Call) Eval(f *fluxer) {
	// fun := reflect.ValueOf(fmt.Println)
	// args := []reflect.Value{}
}
