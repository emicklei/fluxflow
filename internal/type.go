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

func (s TypeSpec) Instantiate(vm *VM) reflect.Value {
	actualType := vm.ReturnsEval(s.Type).Interface()
	// fmt.Println(actualType)
	if i, ok := actualType.(CanInstantiate); ok {
		instance := i.Instantiate(vm)
		// fmt.Println(instance)
		return instance
	}
	panic(fmt.Sprintf("expected a CanInstantiate value:%v", s.Type))
}

func (s TypeSpec) LiteralCompose(composite reflect.Value, values []reflect.Value) reflect.Value {
	if c, ok := s.Type.(CanCompose); ok {
		return c.LiteralCompose(composite, values)
	}
	return expected(s.Type, "a CanCompose value")
}

type StructType struct {
	*ast.StructType
	Fields  *FieldList
	Methods map[string]FuncDecl
}

func makeStructType(ast *ast.StructType) StructType {
	return StructType{StructType: ast,
		Methods: map[string]FuncDecl{},
	}
}

func (s StructType) String() string {
	return fmt.Sprintf("StructType(%v)", s.Fields)
}

func (s StructType) Eval(vm *VM) {
	vm.Returns(reflect.ValueOf(s))
}

func (s StructType) LiteralCompose(composite reflect.Value, values []reflect.Value) reflect.Value {
	i, ok := composite.Interface().(CanCompose)
	if !ok {
		expected(composite, "CanCompose")
	}
	return i.LiteralCompose(composite, values)
}

func (s StructType) Instantiate(vm *VM) reflect.Value {
	return reflect.ValueOf(NewInstance(vm, s))
}
