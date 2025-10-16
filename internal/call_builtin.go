package internal

import "reflect"

// https://pkg.go.dev/builtin#append
func (c CallExpr) evalAppend(vm *VM) {
	args := make([]reflect.Value, len(c.Args))
	for i, arg := range c.Args {
		args[i] = vm.returnsEval(arg)
	}
	result := reflect.Append(args[0], args[1:]...)
	vm.pushOperand(result)
}

// https://pkg.go.dev/builtin#len
func (c CallExpr) evalLen(vm *VM) {
	var sized reflect.Value
	if vm.isStepping {
		sized = vm.callStack.top().pop()
	} else {
		sized = vm.returnsEval(c.Args[0])
	}
	vm.pushOperand(reflect.ValueOf(sized.Len()))
}

// https://pkg.go.dev/builtin#cap
func (c CallExpr) evalCap(vm *VM) {
	var sized reflect.Value
	if vm.isStepping {
		sized = vm.callStack.top().pop()
	} else {
		sized = vm.returnsEval(c.Args[0])
	}
	vm.pushOperand(reflect.ValueOf(sized.Cap()))
}

// https://pkg.go.dev/builtin#clear
// It returns the cleared map or slice.
func (c CallExpr) evalClear(vm *VM) reflect.Value {
	var mapOrSlice reflect.Value
	if vm.isStepping {
		mapOrSlice = vm.callStack.top().pop()
	} else {
		mapOrSlice = vm.returnsEval(c.Args[0])
	}
	mapOrSlice.Clear()
	return mapOrSlice
}
