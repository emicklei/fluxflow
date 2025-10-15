package internal

import (
	"errors"
	"fmt"
	"go/token"

	"golang.org/x/tools/go/packages"
)

type Program struct {
	builder builder // we only need the builder's env TODO
}

func LoadPackages(absolutePath string, optionalConfig *packages.Config) ([]*packages.Package, error) {
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
		return pkgs, fmt.Errorf("failed to load package: %v", err)
	}
	if count := packages.PrintErrors(pkgs); count > 0 {
		return pkgs, fmt.Errorf("errors during package loading: %d", count)
	}
	if len(pkgs) == 0 {
		return pkgs, fmt.Errorf("no packages found")
	}
	return pkgs, nil
}

func BuildProgram(pkgs []*packages.Package, isStepping bool) (*Program, error) {
	b := newBuilder()
	b.opts = buildOptions{callGraph: isStepping}
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

	// run it step by step
	vm.takeAll(decl.callGraph)
	return nil
}
