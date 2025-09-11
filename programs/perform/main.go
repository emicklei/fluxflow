package main

import (
	"flag"
	"go/token"
	"log"

	"github.com/emicklei/fluxflow"
	"golang.org/x/tools/go/packages"
)

var programPath = flag.String("program", "../test1", "Path to the Go program to perform")

func main() {
	flag.Parse()
	fset := token.NewFileSet()
	dirPath := *programPath

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedFiles,
		Fset: fset,
		Dir:  dirPath,
	}
	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		log.Fatalf("failed to load packages: %v", err)
	}
	if packages.PrintErrors(pkgs) > 0 {
		log.Fatal("errors during package loading")
	}

	if len(pkgs) == 0 {
		log.Fatal("no packages found")
	}

	for _, pkg := range pkgs {
		for _, stx := range pkg.Syntax {
			obj := stx.Scope.Lookup("main")
			if obj != nil {
				fluxflow.DoDecl(obj.Decl)
			}
		}
	}
}
