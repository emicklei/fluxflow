package fluxflow

import "github.com/emicklei/fluxflow/internal"

func Run(absolutePath string) error {
	pkgs, err := internal.LoadPackages(absolutePath, nil)
	if err != nil {
		return err
	}
	prog, err := internal.BuildProgram(pkgs, false)
	if err != nil {
		return err
	}
	return internal.RunProgram(prog, nil)
}

/**
type Package = internal.Package

func LoadPackage(absolutePath string) (*Package, error) {
	return nil, nil
}

func Call(pkg *Package, funcName string, args []any) ([]any, error) {
	return nil, nil
}
**/
