package internal

import (
	"fmt"
	"go/token"
	"os"
	"reflect"
	"testing"

	"golang.org/x/tools/go/packages"
)

func loadAndRun(t *testing.T, dirPath string) {
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

	b := builder{}
	for _, pkg := range pkgs {
		for _, stx := range pkg.Syntax {
			for _, decl := range stx.Decls {
				b.Visit(decl)
			}
		}
	}
	runWithBuilder(b)
}
func runWithBuilder(b builder) {
	first := b.first()
	if first == nil {
		return
	}
	here := first
	vm := newVM()
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
		fmt.Print(params...)

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
		fmt.Println(params...)
	}))
	for here != nil {
		hereVal := here.Eval(vm)
		if hereVal.IsValid() && hereVal.CanInterface() {
			fmt.Fprintln(os.Stderr, here, "->", hereVal.Interface())
		} else {
			fmt.Fprintln(os.Stderr, here, "->", hereVal)
		}
		here = here.next
	}
}
