package internal

import "reflect"

type Evaluable interface {
	Eval(vm *VM)
}

type CanAssign interface {
	Assign(env *Env, value reflect.Value)
}

type Statement interface{}

type Expr interface {
	Eval(vm *VM)
}

// All statement nodes implement the Stmt interface.
type Stmt interface {
	stmtStep() Evaluable
}
