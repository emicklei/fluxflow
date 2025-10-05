package fluxflow

import "github.com/emicklei/fluxflow/internal"

func Run(absolutePath string) error {
	prog, err := internal.LoadProgram(absolutePath, nil)
	if err != nil {
		return err
	}
	return internal.RunProgram(prog, nil)
}
