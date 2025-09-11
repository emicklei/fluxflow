package internal

import "reflect"

type FieldSelectable interface {
	Select(name string) reflect.Value
}

type Package struct {
	Name string
}

func (p Package) Select(field string) reflect.Value {
	return stdpkg[p.Name][field]
}
