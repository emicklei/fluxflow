package main

import (
	"fmt"

	"github.com/emicklei/fluxflow/examples/subpkg/pkg"
)

func main() {
	fmt.Println(pkg.Name, pkg.IsWeekend("Sunday"))
}
