package internal

import (
	"fmt"
	"go/token"
	"io"
	"os"
	"os/exec"
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

func parseAndStepThrough(t *testing.T, source string) string {
	prog := buildProgram(t, source)
	vm := newVM(prog.builder.env)
	collectPrintOutput(vm)

	main := prog.builder.env.valueLookUp("main")
	decl := main.Interface().(FuncDecl)

	g := new(grapher)
	decl.Flow(g)
	g.dotify()
	// will fail in pipeline without graphviz installed
	exec.Command("dot", "-Tpng", "-o", "graph.png", "graph.dot").Run()

	// run it step by step
	vm.isStepping = true
	here := g.head
	for here != nil {
		if trace {
			fmt.Println("taking", here)
		}
		here = here.Take(vm)
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
