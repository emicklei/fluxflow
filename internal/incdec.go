package internal

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type IncDecStmt struct {
	*ast.IncDecStmt
	X Expr
}

func (i IncDecStmt) stmtStep() Evaluable { return i }

func (i IncDecStmt) String() string {
	return fmt.Sprintf("IncDecStmt(%v)", i.X)
}
func (i IncDecStmt) Eval(vm *VM) {
	current := vm.ReturnsEval(i.X)
	if i.Tok == token.INC {
		switch current.Kind() {
		case reflect.Int, reflect.Int64:
			if a, ok := i.X.(CanAssign); ok {
				a.Assign(vm.localEnv(), reflect.ValueOf(current.Int()+1))
			}
		case reflect.Float64:
		default:
			panic("unsupported type for ++: " + current.Kind().String())
		}
	} else { // DEC
		switch current.Kind() {
		case reflect.Int, reflect.Int64:
			if a, ok := i.X.(CanAssign); ok {
				a.Assign(vm.localEnv(), reflect.ValueOf(current.Int()-1))
			}
		case reflect.Float64:
		default:
			panic("unsupported type for -- :" + current.Kind().String())
		}
	}
}
