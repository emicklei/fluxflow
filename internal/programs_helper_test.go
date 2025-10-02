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
	b := newBuilder()
	for _, pkg := range pkgs {
		for _, stx := range pkg.Syntax {
			for _, decl := range stx.Decls {
				b.Visit(decl)
			}
		}
	}
	return runWithBuilder(b)
}

func printSteps() func() {
	os.Setenv("STEPS", "1")
	return func() {
		os.Unsetenv("STEPS")
	}
}

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
	builtins := &Environment{valueTable: builtinsMap}
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
	vm := newVM(b.env)
	vm.localEnv().set("print", reflect.ValueOf(func(args ...any) {
		for _, a := range args {
			if rv, ok := a.(reflect.Value); ok && rv.IsValid() && rv.CanInterface() {
				fmt.Fprintf(vm.output, "%v", rv.Interface())
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
	// first run const and vars
	// try declare all of them until none left
	// a declare may refer to other unseen declares.
	pkgEnv := vm.localEnv().(*PkgEnvironment)
	for len(pkgEnv.declTable) > 0 {
		for key, each := range pkgEnv.declTable {
			if each.Declare(vm) {
				delete(pkgEnv.declTable, key)
			}
		}
	}
	// TODO first run inits

	main := vm.localEnv().valueLookUp("main")
	if !main.IsValid() {
		return "main not found"
	}
	// TODO
	vm.callStack.push(stackFrame{env: vm.localEnv().newChildEnvironment()})
	fundecl := main.Interface().(FuncDecl)
	fundecl.Body.Eval(vm)
	return vm.output.String()
}
