package internal

import (
	"fmt"
	"go/ast"
)

type FuncDecl struct {
	*ast.FuncDecl
	Name      *Ident
	Recv      *FieldList
	Body      *BlockStmt
	Type      *FuncType
	callGraph Step
}

func (f FuncDecl) Eval(vm *VM) {
	if f.Body != nil {
		vm.eval(f.Body)
	}
}

func (f FuncDecl) Flow(g *grapher) (head Step) {
	head = g.current
	if f.Body != nil {
		head = f.Body.Flow(g)
	}
	return
}

// Take the function call graph and execute it step by step
func (f FuncDecl) Take(vm *VM) Step {
	here := f.callGraph
	for here != nil {
		if trace {
			fmt.Println("funcdecl taking", here)
		}
		here = here.Take(vm)
	}
	return nil
}

func (f FuncDecl) String() string {
	return fmt.Sprintf("FuncDecl(%s)", f.Name.Name)
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
