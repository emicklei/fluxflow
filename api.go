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
