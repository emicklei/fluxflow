package internal

import (
	"fmt"
	"go/ast"
	"log"
	"reflect"
)

// corresponding Field, XxxSpec, FuncDecl, LabeledStmt, AssignStmt, Scope; or nil
func BuildSteps(decl any) any {
	f, ok := decl.(*ast.FuncDecl)
	if !ok {
		log.Println(decl)
		return nil
	}
	b := builder{}
	b.Visit(f.Body)
	return b.first()
}
func RunSteps(decl any) {
	first := BuildSteps(decl)
	if first == nil {
		return
	}
	here := first.(*step)
	vm := newVM()
	// builtin
	vm.env.set("print", reflect.ValueOf(func(args ...any) {
		for _, a := range args {
			if rv, ok := a.(reflect.Value); ok && rv.IsValid() && rv.CanInterface() {
				fmt.Print(rv.Elem())
			} else {
				fmt.Print(a)
			}
		}
	}))
	for here != nil {
		hereVal := here.Eval(vm)
		if hereVal.IsValid() && hereVal.CanInterface() {
			fmt.Println(here, "->", hereVal.Interface())
		} else {
			fmt.Println(here, "->", hereVal)
		}
		here = here.next
	}
}
