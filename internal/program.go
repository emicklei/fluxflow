package internal

import (
	"errors"
	"fmt"
	"go/token"
	"os"
	"os/exec"

	"golang.org/x/tools/go/packages"
)

type Program struct {
	builder builder
}

func LoadProgram(absolutePath string, optionalConfig *packages.Config) (*Program, error) {
	var cfg *packages.Config
	if optionalConfig != nil {
		cfg = optionalConfig
	} else {
		cfg = &packages.Config{
			Mode: packages.NeedName | packages.NeedSyntax | packages.NeedFiles,
			Fset: token.NewFileSet(),
			Dir:  absolutePath,
		}
	}
	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		return nil, fmt.Errorf("failed to load package: %v", err)
	}
	if count := packages.PrintErrors(pkgs); count > 0 {
		return nil, fmt.Errorf("errors during package loading: %d", count)
	}
	if len(pkgs) == 0 {
		return nil, fmt.Errorf("no packages found")
	}
	b := newBuilder()
	for _, pkg := range pkgs {
		for _, stx := range pkg.Syntax {
			for _, decl := range stx.Decls {
				b.Visit(decl)
			}
		}
	}
	return &Program{builder: b}, nil
}

func RunProgram(p *Program, optionalVM *VM) error {
	var vm *VM
	if optionalVM != nil {
		vm = optionalVM
	} else {
		vm = newVM(p.builder.env)
	}
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
	// then run all inits
	for _, each := range pkgEnv.inits {
		vm.pushNewFrame()
		each.Eval(vm)
		vm.popFrame()
	}
	main := vm.localEnv().valueLookUp("main")
	if !main.IsValid() {
		return errors.New("main not found")
	}
	// TODO
	vm.pushNewFrame()
	fundecl := main.Interface().(FuncDecl)
	fundecl.Eval(vm)
	vm.popFrame()
	return nil
}

func WalkProgram(p *Program, optionalVM *VM) error {
	var vm *VM
	if optionalVM != nil {
		vm = optionalVM
	} else {
		vm = newVM(p.builder.env)
	}
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

	// then walk all inits
	// TODO should walk through them step by step!!
	for _, each := range pkgEnv.inits {
		vm.pushNewFrame()
		each.Eval(vm)
		vm.popFrame()
	}

	main := p.builder.env.valueLookUp("main")
	if !main.IsValid() {
		return errors.New("main not found")
	}
	decl := main.Interface().(FuncDecl)

	idgen = 0 // state of grapher?
	g := new(grapher)
	decl.Flow(g)
	if fileName := os.Getenv("DOT"); fileName != "" {
		g.dotFile = fileName
		g.dotify()
		// will fail in pipeline without graphviz installed
		exec.Command("dot", "-Tpng", "-o", g.dotFilename()+".png", g.dotFilename()).Run()
	}

	// run it step by step
	vm.isStepping = true
	here := g.head
	for here != nil {
		if trace {
			fmt.Println("taking", here)
		}
		here = here.Take(vm)
	}
	return nil
}
