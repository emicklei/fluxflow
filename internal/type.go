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
	fmt.Println(actualType)
	if i, ok := actualType.(CanInstantiate); ok {
		instance := i.Instantiate(vm)
		fmt.Println(instance)
		return instance
	}
	panic(fmt.Sprintf("expected a CanInstantiate value:%v", s.Type))
}

func (s TypeSpec) LiteralCompose(composite reflect.Value, values []reflect.Value) reflect.Value {
	if c, ok := s.Type.(CanCompose); ok {
		return c.LiteralCompose(composite, values)
	}
	panic(fmt.Sprintf("expected a CanCompose value:%v", s.Type))
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

// func (s StructType) LiteralCompose(composite reflect.Value, values []reflect.Value) reflect.Value {
// 	return reflect.ValueOf(NewInstance(s))
// }

func (s StructType) Instantiate(vm *VM) reflect.Value {
	return reflect.ValueOf(NewInstance(vm, s))
}

// first for struct
type Instance struct {
	Type   StructType
	fields map[string]reflect.Value
}

func NewInstance(vm *VM, t StructType) Instance {
	i := Instance{Type: t, fields: map[string]reflect.Value{}}
	for _, field := range t.Fields.List {
		for _, name := range field.Names {
			i.fields[name.Name] = reflect.Value{} // field.Type.ZeroValue(vm.localEnv())
		}
	}
	return i
}
func (i Instance) String() string {
	return fmt.Sprintf("Instance(%v)", i.Type)
}

func (i Instance) Select(name string) reflect.Value {
	if v, ok := i.fields[name]; ok {
		return v
	}
	panic("no such field: " + name)
}
func (i Instance) LiteralCompose(values []reflect.Value) reflect.Value {
	return reflect.ValueOf("todo")
}
