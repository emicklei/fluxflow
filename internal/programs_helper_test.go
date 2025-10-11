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
	cwd, _ := os.Getwd()
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedFiles,
		Fset: token.NewFileSet(),
		Dir:  path.Join(cwd, "../examples"),
		Overlay: map[string][]byte{
			path.Join(cwd, "../examples/main.go"): []byte(source),
		},
	}
	prog, err := LoadProgram(cfg.Dir, cfg)
	if err != nil {
		t.Fatalf("failed to load package: %v", err)
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
	prog := buildProgram(t, source)
	vm := newVM(prog.builder.env)
	collectPrintOutput(vm)
	if err := WalkProgram(prog, vm); err != nil {
		panic(err)
	}
	return vm.output.String()
}

func parseAndRun(t *testing.T, source string) string {
	prog := buildProgram(t, source)
	vm := newVM(prog.builder.env)
	collectPrintOutput(vm)
	if err := RunProgram(prog, vm); err != nil {
		panic(err)
	}
	return vm.output.String()
}
