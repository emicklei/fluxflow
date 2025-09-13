package internal

import (
	"go/token"
	"log"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestProgram0(t *testing.T) {
	loadAndRun(t, "../programs/test0", func(obj any) {
		RunSteps(obj)
	})
}

func TestProgram1(t *testing.T) {
	loadAndRun(t, "../programs/test1", func(obj any) {
		BuildSteps(obj)
		//s := DoDecl(obj)
		//structexplorer.NewService("some structure", s).Start()
	})
}

func TestProgram2(t *testing.T) {
	loadAndRun(t, "../programs/test2", func(obj any) {
		BuildSteps(obj)
	})
}
func TestProgram3(t *testing.T) {
	loadAndRun(t, "../programs/test3", func(obj any) {
		// Show(DoDecl(obj))
		BuildSteps(obj)
	})
}
func loadAndRun(t *testing.T, dirPath string, fn func(obj any)) {
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

	for _, pkg := range pkgs {
		for _, stx := range pkg.Syntax {
			obj := stx.Scope.Lookup("main")
			if obj != nil {
				fn(obj.Decl)
			} else {
				log.Printf("no main function found")
			}
		}
	}
}
