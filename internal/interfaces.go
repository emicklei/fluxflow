package internal

import (
	"reflect"

	"github.com/emicklei/dot"
)

type Evaluable interface {
	Eval(vm *VM)
}

type CanAssign interface {
	Assign(vm *VM, value reflect.Value)
	Define(vm *VM, value reflect.Value) // needed?
}

type CanDeclare interface {
	Declare(vm *VM) bool
}

type Statement interface{}

type Expr interface {
	//Flowable
	Evaluable
}

type Flowable interface {
	Flow(g *grapher)
}

type HasZeroValue interface {
	ZeroValue(env Env) reflect.Value
}

// All statement nodes implement the Stmt interface.
type Stmt interface {
	Flowable
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

type Decl interface {
	declStep() CanDeclare
}

type Step interface {
	Evaluable
	SetNext(s Step)
	Next() Step
	ID() int
	String() string
	Traverse(g *dot.Graph, visited map[int]dot.Node) dot.Node
}
