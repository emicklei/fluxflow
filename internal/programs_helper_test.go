package internal

import (
	"fmt"
	"go/token"
	"io"
	"os"
	"path"
	"reflect"
	"testing"

	"golang.org/x/tools/go/packages"
)

func parseAndRun(t *testing.T, source string) string {
	fset := token.NewFileSet()

	cwd, _ := os.Getwd()
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedFiles,
		Fset: fset,
		Dir:  path.Join(cwd, "../programs"),
		Overlay: map[string][]byte{
			path.Join(cwd, "../programs/main.go"): []byte(source),
		},
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
	builtins := &Env{symbolTable: builtinsMap}
	b := builder{env: builtins}
	for _, pkg := range pkgs {
		for _, stx := range pkg.Syntax {
			for _, decl := range stx.Decls {
				b.Visit(decl)
			}
		}
	}
	return runWithBuilder(b)
}

func printSteps() {
	os.Setenv("STEPS", "1")
}

func loadAndRun(t *testing.T, dirPath string) string {
	fset := token.NewFileSet()

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedFiles,
		Fset: fset,
		Dir:  dirPath,
		Overlay: map[string][]byte{
			"/Users/ernestmicklei/Projects/fluxflow/programs/test_print/main.go": []byte(`package main`),
		},
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
	builtins := &Env{symbolTable: builtinsMap}
	b := builder{env: builtins}
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
	vm.env.set("print", reflect.ValueOf(func(args ...any) {
		for _, a := range args {
			if rv, ok := a.(reflect.Value); ok && rv.IsValid() && rv.CanInterface() {
				fmt.Fprintf(vm.output, "%v", rv.Elem())
			} else {
				if s, ok := a.(string); ok {
					io.WriteString(vm.output, s)
					continue
				} else {
					fmt.Fprintf(vm.output, "%v", a)
				}
			}
		}
	}))
	main := vm.env.lookUp("main")
	if !main.IsValid() {
		return "main not found"
	}
	// TODO
	vm.callStack.push(stackFrame{env: vm.env.subEnv()})
	fundecl := main.Interface().(FuncDecl)
	fundecl.Body.Eval(vm)
	return vm.output.String()
}
