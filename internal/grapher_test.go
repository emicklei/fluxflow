package internal

import (
	"testing"
)

func TestGrapherFor(t *testing.T) {
	//t.Skip()
	source := `
package main

func main() {
	j := 1
	for i := 0; i < 3; i++ {
		j = i
		print(i)
	}
	print(j)
}`

	out := parseAndStepThrough(t, source)
	/**
		prog := buildProgram(t, source)
		main := prog.builder.env.valueLookUp("main")
		decl := main.Interface().(FuncDecl)

		g := new(grapher)
		decl.Flow(g)
		g.dotify()
		// will fail in pipeline without graphviz installed
		exec.Command("dot", "-Tpng", "-o", "graph.png", "graph.dot").Run()

		// run it step by step
		vm := newVM(prog.builder.env)
		collectPrintOutput(vm)
		vm.isStepping = true
		here := g.head
		for here != nil {
			fmt.Println(here)
			here = here.Take(vm)
		}
		t.Log(vm.output.String())
	**/
	expected := `0122`
	if out != expected {
		t.Fatalf("expected %q got %q", expected, out)
	}
}
