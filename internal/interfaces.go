package internal

import "reflect"

type Evaluable interface {
	Eval(vm *VM) reflect.Value
}

type CanAssign interface {
	Assign(env *Env, value reflect.Value)
}

type Statement interface{}

type Expr interface {
	Eval(vm *VM) reflect.Value
}
