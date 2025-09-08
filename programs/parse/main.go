package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"

	"golang.org/x/tools/go/packages"
)

func main() {
	fset := token.NewFileSet()
	dirPath := "../test2"

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
		fmt.Println("Package:", pkg.Name)
		for _, f := range pkg.Syntax {
			// ast.Print writes to os.Stdout
			err := ast.Print(fset, f)
			if err != nil {
				log.Fatalf("failed to print ast for file: %v", err)
			}
		}
	}
}
