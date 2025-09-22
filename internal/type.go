package internal

import (
	"fmt"
	"go/ast"
	"reflect"
)

type TypeSpec struct {
	*ast.TypeSpec
	Name       *Ident
	TypeParams *FieldList
	Type       Expr
}

func (s TypeSpec) String() string {
	return fmt.Sprintf("TypeSpec(%v,%v,%v)", s.Name, s.TypeParams, s.Type)
}

func (s TypeSpec) Eval(vm *VM) {
	if s.Name == nil {
		return // TODO ?
	}
	actualType := vm.ReturnsEval(s.Type)
	vm.localEnv().set(s.Name.Name, actualType) // use the spec itself as value
}

func (s TypeSpec) LiteralCompose(comp reflect.Value, vals reflect.Value) reflect.Value {
	return reflect.Value{}
}

type StructType struct {
	*ast.StructType
	Fields *FieldList
}

func (s StructType) String() string {
	return fmt.Sprintf("StructType(%v)", s.Fields)
}

func (s StructType) Eval(vm *VM) {
	vm.Returns(reflect.ValueOf(s))
}

// first for struct
type Instance struct {
	Type StructType
}
