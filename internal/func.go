package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type FuncDecl struct {
	Name *Ident
	Body *BlockStmt
	*ast.FuncDecl
}

func (f *FuncDecl) Eval(vm *VM) reflect.Value {
	return reflect.Value{}
}
func (f *FuncDecl) String() string {
	return fmt.Sprintf("FuncDecl(%s)", f.Name.Name)
}
