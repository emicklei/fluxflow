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

func buildProgram(t *testing.T, source string) *Program {
	t.Helper()
	cwd, _ := os.Getwd()
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedFiles,
		Fset: token.NewFileSet(),
		Dir:  path.Join(cwd, "../examples"),
		Overlay: map[string][]byte{
			path.Join(cwd, "../examples/main.go"): []byte(source),
		},
	}
	pkgs, err := LoadPackages(cfg.Dir, cfg)
	if err != nil {
		t.Fatalf("failed to load packages: %v", err)
	}
	prog, err := BuildProgram(pkgs, true)
	if err != nil {
		t.Fatalf("failed to build program: %v", err)
	}
	return prog
}

func collectPrintOutput(vm *VM) {
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
}

func parseAndWalk(t *testing.T, source string) string {
	t.Helper()
	idgen = 0
	prog := buildProgram(t, source)
	vm := newVM(prog.builder.env)
	collectPrintOutput(vm)
	if err := WalkProgram(prog, vm); err != nil {
		panic(err)
	}
	return vm.output.String()
}

func parseAndRun(t *testing.T, source string) string {
	t.Helper()
	prog := buildProgram(t, source)
	vm := newVM(prog.builder.env)
	collectPrintOutput(vm)
	if err := RunProgram(prog, vm); err != nil {
		t.Fatal(err)
	}
	return vm.output.String()
}

func testProgram(t *testing.T, running bool, stepping bool, source string, wantFuncOrString any) {
	t.Parallel()
	t.Helper()
	if running {
		out := parseAndRun(t, source)
		if fn, ok := wantFuncOrString.(func(string) bool); ok {
			if !fn(out) {
				t.Errorf("got [%v] which does not match predicate", out)
			}
			return
		}
		want := wantFuncOrString.(string)
		if got, want := out, want; got != want {
			t.Errorf("[run] got [%v] want [%v]", got, want)
		}
	} else {
		t.Log("TODO skipping running through:", t.Name())
	}
	if stepping {
		os.WriteFile(fmt.Sprintf("testgraphs/%s.src", t.Name()), []byte(source), 0644)
		os.Setenv("DOT", fmt.Sprintf("testgraphs/%s.dot", t.Name()))
		out := parseAndWalk(t, source)
		os.Unsetenv("DOT")
		if fn, ok := wantFuncOrString.(func(string) bool); ok {
			if !fn(out) {
				t.Errorf("got [%v] which does not match predicate", out)
			}
			return
		}
		want := wantFuncOrString.(string)
		if got, want := out, want; got != want {
			t.Errorf("[step] got [%v] want [%v]", got, want)
		}
	} else {
		t.Log("TODO skipping stepping through:", t.Name())
	}
}
