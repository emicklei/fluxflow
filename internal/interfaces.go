package internal

import "reflect"

type Evaluable interface {
	Eval(vm *VM)
}

type CanAssign interface {
	Assign(env ienv, value reflect.Value)
	Define(env ienv, value reflect.Value) // needed?
}

type CanDefine interface {
	Define(env ienv, value reflect.Value)
}

type CanDeclare interface {
	Declare(vm *VM)
}

type Statement interface{}

type Expr interface {
	Eval(vm *VM)
}

type HasZeroValue interface {
	ZeroValue(env ienv) reflect.Value
}

// All statement nodes implement the Stmt interface.
type Stmt interface {
	stmtStep() Evaluable
}

type CanCompose interface {
	LiteralCompose(composite reflect.Value, values []reflect.Value) reflect.Value
}

type FieldSelectable interface {
	Select(name string) reflect.Value
}

type CanInstantiate interface {
	Instantiate(vm *VM) reflect.Value // , typeArgs []reflect.Type) reflect.Value
	LiteralCompose(composite reflect.Value, values []reflect.Value) reflect.Value
}
