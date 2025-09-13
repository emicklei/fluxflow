package fluxflow

import (
	"github.com/emicklei/fluxflow/internal"
)

func DoDecl(decl any) {
	internal.BuildSteps(decl)
}
