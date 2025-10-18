package internal

import (
	"fmt"
	"reflect"
)

type Package struct {
	Name        string // key in env
	Path        string
	Env         Env
	Initialized bool
}

func (p Package) String() string {
	return fmt.Sprintf("Package(%s,%s)", p.Name, p.Path)
}
func (p Package) Select(name string) reflect.Value {
	symbolTable, ok := stdpkg[p.Path]
	if !ok {
		return reflect.Value{}
	}
	v, ok := symbolTable[name]
	if !ok {
		return reflect.Value{}
	}
	return v
}
