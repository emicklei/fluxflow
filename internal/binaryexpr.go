package internal

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type BinaryExpr struct {
	operatorUnimplemented
	X Expr
	Y Expr
	*ast.BinaryExpr
}

func (s BinaryExpr) Eval(vm *VM) {
	v := BinaryExprValue{
		left:  vm.returnsEval(s.X),
		op:    s.Op,
		right: vm.returnsEval(s.Y),
	}
	// TODO
	if !v.IsValid() {
		vm.pushOperand(reflect.Value{})
		return
	}
	vm.pushOperand(v.Eval())
}

func (s BinaryExpr) String() string {
	return fmt.Sprintf("BinaryExpr(%v %v %v)", s.X, s.Op, s.Y)
}

type BinaryExprValue struct {
	left  reflect.Value
	op    token.Token
	right reflect.Value
}

func (b BinaryExprValue) IsValid() bool {
	return b.left.IsValid() && b.right.IsValid()
}

func (b BinaryExprValue) Eval() reflect.Value {
	switch b.left.Kind() {
	case reflect.Int:
		res := b.IntEval(b.left.Int())
		if res.CanInt() {
			return reflect.ValueOf(int(res.Int()))
		} else {
			return res
		}
	case reflect.Uint:
		res := b.UIntEval(b.left.Uint())
		if res.CanInt() {
			return reflect.ValueOf(uint(res.Int()))
		} else {
			return res
		}
	case reflect.Int8:
		res := b.IntEval(b.left.Int())
		if res.CanInt() {
			return reflect.ValueOf(int8(res.Int()))
		} else {
			return res
		}
	case reflect.Uint8:
		res := b.UIntEval(b.left.Uint())
		if res.CanInt() {
			return reflect.ValueOf(uint8(res.Int()))
		} else {
			return res
		}
	case reflect.Int16:
		res := b.IntEval(b.left.Int())
		if res.CanInt() {
			return reflect.ValueOf(int16(res.Int()))
		} else {
			return res
		}
	case reflect.Uint16:
		res := b.UIntEval(b.left.Uint())
		if res.CanInt() {
			return reflect.ValueOf(uint16(res.Int()))
		} else {
			return res
		}
	case reflect.Int32:
		res := b.IntEval(b.left.Int())
		if res.CanInt() {
			return reflect.ValueOf(int32(res.Int()))
		} else {
			return res
		}
	case reflect.Uint32:
		res := b.UIntEval(b.left.Uint())
		if res.CanInt() {
			return reflect.ValueOf(uint32(res.Int()))
		} else {
			return res
		}
	case reflect.Int64:
		return b.IntEval(b.left.Int())
	case reflect.Uint64:
		return reflect.ValueOf(b.UIntEval(b.left.Uint()).Uint())
	// non-ints
	case reflect.Float64:
		return b.FloatEval(b.left.Float())
	case reflect.String:
		if b.op == token.ADD && b.right.Kind() == reflect.String {
			return reflect.ValueOf(b.left.String() + b.right.String())
		}
	}
	panic("not implemented: BinaryExprValue.Eval:" + b.left.Kind().String())
}
func (b BinaryExprValue) IntEval(left int64) reflect.Value {
	switch b.right.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return b.IntOpInt(left, b.right.Int())
	case reflect.Float64:
		return b.FloatOpFloat(float64(left), b.right.Float())
	}
	panic("not implemented: BinaryExprValue.IntEval:" + b.right.Kind().String())
}

func (b BinaryExprValue) UIntEval(left uint64) reflect.Value {
	switch b.right.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return b.UIntOpUInt(left, b.right.Uint())
	}
	panic("not implemented: BinaryExprValue.UIntEval:" + b.right.Kind().String())
}

func (b BinaryExprValue) FloatEval(left float64) reflect.Value {
	switch b.right.Kind() {
	case reflect.Float64:
		return b.FloatOpFloat(left, b.right.Float())
	case reflect.Int:
		return b.FloatOpFloat(left, float64(b.right.Int()))
	}
	panic("not implemented: BinaryExprValue.FloatEval:" + b.right.Kind().String())
}

func (b BinaryExprValue) FloatOpFloat(left float64, right float64) reflect.Value {
	switch b.op {
	case token.ADD:
		return reflect.ValueOf(left + right)
	case token.SUB:
		return reflect.ValueOf(left - right)
	case token.MUL:
		return reflect.ValueOf(left * right)
	case token.QUO:
		return reflect.ValueOf(left / right)
	}
	panic("not implemented: BinaryExprValue.FloatOpFloat:" + b.op.String())
}

func (b BinaryExprValue) IntOpInt(left int64, right int64) reflect.Value {
	switch b.op {
	case token.ADD:
		return reflect.ValueOf(left + right)
	case token.SUB:
		return reflect.ValueOf(left - right)
	case token.MUL:
		return reflect.ValueOf(left * right)
	case token.QUO:
		return reflect.ValueOf(left / right)
	case token.REM:
		return reflect.ValueOf(left % right)
	case token.AND:
		return reflect.ValueOf(left & right)
	case token.OR:
		return reflect.ValueOf(left | right)
	case token.XOR:
		return reflect.ValueOf(left ^ right)
	case token.SHL:
		// right must be unsigned
		return reflect.ValueOf(left << uint64(right))
	case token.SHR:
		// right must be unsigned
		return reflect.ValueOf(left >> uint64(right))
	case token.AND_NOT:
		return reflect.ValueOf(left &^ right)
	case token.EQL:
		return reflect.ValueOf(left == right)
	case token.NEQ:
		return reflect.ValueOf(left != right)
	case token.LSS:
		return reflect.ValueOf(left < right)
	case token.LEQ:
		return reflect.ValueOf(left <= right)
	case token.GTR:
		return reflect.ValueOf(left > right)
	case token.GEQ:
		return reflect.ValueOf(left >= right)
	}
	panic("not implemented: BinaryExprValue.IntOpInt:" + b.op.String())
}

func (b BinaryExprValue) UIntOpUInt(left uint64, right uint64) reflect.Value {
	switch b.op {
	case token.ADD:
		return reflect.ValueOf(left + right)
	case token.SUB:
		return reflect.ValueOf(left - right)
	case token.MUL:
		return reflect.ValueOf(left * right)
	case token.QUO:
		return reflect.ValueOf(left / right)
	case token.REM:
		return reflect.ValueOf(left % right)
	case token.AND:
		return reflect.ValueOf(left & right)
	case token.OR:
		return reflect.ValueOf(left | right)
	case token.XOR:
		return reflect.ValueOf(left ^ right)
	case token.SHL:
		return reflect.ValueOf(left << right)
	case token.SHR:
		return reflect.ValueOf(left >> right)
	case token.AND_NOT:
		return reflect.ValueOf(left &^ right)
	case token.EQL:
		return reflect.ValueOf(left == right)
	case token.NEQ:
		return reflect.ValueOf(left != right)
	case token.LSS:
		return reflect.ValueOf(left < right)
	case token.LEQ:
		return reflect.ValueOf(left <= right)
	case token.GTR:
		return reflect.ValueOf(left > right)
	case token.GEQ:
		return reflect.ValueOf(left >= right)
	}
	panic("not implemented: BinaryExprValue.UIntOpUInt:" + b.op.String())
}

var _ Expr = &operatorUnimplemented{}

type operatorUnimplemented struct{ step }

func (*operatorUnimplemented) Assign(env *Environment, value reflect.Value) {
	panic("not implemented: operatorUnimplemented.Assign")
}
