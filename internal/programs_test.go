package internal

import (
	"go/token"
	"log"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestProgramPrint(t *testing.T) {
	loadAndRun(t, "../programs/test_print", func(obj any) {
		RunSteps(obj)
	})
}

func TestProgramMulitAssign(t *testing.T) {
	loadAndRun(t, "../programs/test_multiassign", func(obj any) {
		BuildSteps(obj)
		//s := DoDecl(obj)
		//structexplorer.NewService("some structure", s).Start()
	})
}

func TestProgramGeneric(t *testing.T) {
	loadAndRun(t, "../programs/test_generic", func(obj any) {
		BuildSteps(obj)
	})
}
func TestProgramTypeAssert(t *testing.T) {
	loadAndRun(t, "../programs/test_typeassert", func(obj any) {
		// Show(DoDecl(obj))
		BuildSteps(obj)
	})
}

func TestProgramIf(t *testing.T) {
	loadAndRun(t, "../programs/test_if", func(obj any) {
		RunSteps(obj)
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
