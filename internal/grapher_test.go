package internal

import (
	"os/exec"
	"testing"
)

func TestGrapherFor(t *testing.T) {
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
	prog := buildProgram(t, source)
	main := prog.builder.env.valueLookUp("main")
	decl := main.Interface().(FuncDecl)
	forstmt := decl.Body.List[1].(ForStmt)

	g := new(grapher)
	forstmt.Flow(g)
	g.dotify()
	// will fail in pipeline without graphviz installed
	exec.Command("dot", "-Tpng", "-o", "graph.png", "graph.dot").Run()
}
