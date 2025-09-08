package nodes

import (
	"reflect"
	"sort"
)

var stdpkg = map[string]map[string]reflect.Value{}

func init() {
	// This function is intentionally left empty.
	// The stdpkg map is populated by the generated code in stdlib_generated.go.
}

// StandardPackageFunctions returns a list of all known exported function names for a given Go standard library package.
func StandardPackageFunctions(pkgName string) []string {
	if funcs, ok := stdpkg[pkgName]; ok {
		names := make([]string, 0, len(funcs))
		for name := range funcs {
			names = append(names, name)
		}
		// sort for consistent output
		sort.Strings(names)
		return names
	}
	return nil
}
