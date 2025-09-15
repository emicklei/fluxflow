package internal

import (
	"fmt"
	"go/token"
	"reflect"
	"testing"

	"golang.org/x/tools/go/packages"
)

func loadAndRun(t *testing.T, dirPath string) string {
	fset := token.NewFileSet()

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedFiles,
		Fset: fset,
		Dir:  dirPath,
	}
	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		t.Fatalf("failed to load package: %v", err)
	}
	if packages.PrintErrors(pkgs) > 0 {
		t.Fatal("errors during package loading")
	}

	if len(pkgs) == 0 {
		t.Fatal("no packages found")
	}

	b := builder{env: newEnv()}
	for _, pkg := range pkgs {
		for _, stx := range pkg.Syntax {
			for _, decl := range stx.Decls {
				b.Visit(decl)
			}
		}
	}
	return runWithBuilder(b)
}
func runWithBuilder(b builder) string {
	vm := newVM()
	vm.env = b.env
	// builtin
	vm.env.set("print", reflect.ValueOf(func(args ...any) {
		params := make([]any, len(args))
		for _, a := range args {
			if rv, ok := a.(reflect.Value); ok && rv.IsValid() && rv.CanInterface() {
				params = append(params, rv.Elem())
			} else {
				params = append(params, a)
			}
		}
		fmt.Fprint(vm.output, params...)

	}))
	vm.env.set("println", reflect.ValueOf(func(args ...any) {
		params := make([]any, len(args))
		for _, a := range args {
			if rv, ok := a.(reflect.Value); ok && rv.IsValid() && rv.CanInterface() {
				params = append(params, rv.Elem())
			} else {
				params = append(params, a)
			}
		}
		fmt.Fprintln(vm.output, params...)
	}))
	main := vm.env.lookUp("main")
	if !main.IsValid() {
		return "main not found"
	}
	// TODO
	fundecl := main.Interface().(*FuncDecl)
	fundecl.Body.Eval(vm)
	return vm.output.String()
}
