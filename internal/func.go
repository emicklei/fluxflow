package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type FuncDecl struct {
	Name *Ident
	Body *BlockStmt
	Type *FuncType
	*ast.FuncDecl
}

func (f *FuncDecl) Eval(vm *VM) {
	v := reflect.ValueOf(f)
	vm.env.set(f.Name.Name, v)
}
func (f *FuncDecl) String() string {
	return fmt.Sprintf("FuncDecl(%s)", f.Name.Name)
}

type FuncType struct {
	*ast.FuncType
	TypeParams *FieldList
	Params     *FieldList
	Returns    *FieldList
}
