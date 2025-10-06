package internal

import (
	"fmt"
	"go/ast"
)

type FuncDecl struct {
	*ast.FuncDecl
	Name *Ident
	Recv *FieldList
	Body *BlockStmt
	Type *FuncType
	// for handling labeled statements, lazy initialized
	labeledSteps map[string]step
}

func (f FuncDecl) Eval(vm *VM) {
	panic("todo")
}
func (f FuncDecl) String() string {
	return fmt.Sprintf("FuncDecl(%s)", f.Name.Name)
}

func (f FuncDecl) withLabeledStep(label string, s step) FuncDecl {
	if f.labeledSteps == nil {
		f.labeledSteps = map[string]step{label: s}
		return f
	}
	f.labeledSteps[label] = s
	return f
}

type FuncType struct {
	*ast.FuncType
	TypeParams *FieldList
	Params     *FieldList
	Returns    *FieldList
}

func (t FuncType) String() string {
	return fmt.Sprintf("FuncType(%v,%v,%v)", t.TypeParams, t.Params, t.Returns)
}

func (t FuncType) Eval(vm *VM) {}

type Ellipsis struct {
	*ast.Ellipsis
	Elt Expr // ellipsis element type (parameter lists only); or nil
}

func (e Ellipsis) String() string {
	return fmt.Sprintf("Ellipsis(%v)", e.Elt)
}
func (e Ellipsis) Eval(vm *VM) {}
