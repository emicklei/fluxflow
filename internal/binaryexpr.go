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

func (s BinaryExpr) Eval(vm *VM) reflect.Value {
	v := BinaryExprValue{
		left:  s.X.Eval(vm),
		op:    s.Op,
		right: s.Y.Eval(vm),
	}
	return v.Eval()
}

func (s BinaryExpr) String() string {
	return fmt.Sprintf("BinaryExpr(%v %v %v)", s.X, s.Op, s.Y)
}

type BinaryExprValue struct {
	left  reflect.Value
	op    token.Token
	right reflect.Value
}

func (b BinaryExprValue) Eval() reflect.Value {
	switch b.left.Kind() {
	case reflect.Int:
		return b.IntEval(b.left.Int())
	case reflect.Float64:
		return b.FloatEval(b.left.Float())
	case reflect.String:
		if b.op == token.ADD && b.right.Kind() == reflect.String {
			return reflect.ValueOf(b.left.String() + b.right.String())
		}
	}
	panic("not implemented:" + b.left.Kind().String())
}
func (b BinaryExprValue) IntEval(left int64) reflect.Value {
	switch b.right.Kind() {
	case reflect.Int:
		return b.IntOpInt(left, b.right.Int())
	case reflect.Float64:
		return b.FloatOpFloat(float64(left), b.right.Float())
	}
	panic("not implemented:" + b.right.Kind().String())
}

func (b BinaryExprValue) FloatEval(left float64) reflect.Value {
	switch b.right.Kind() {
	case reflect.Float64:
		return b.FloatOpFloat(left, b.right.Float())
	case reflect.Int:
		return b.FloatOpFloat(left, float64(b.right.Int()))
	}
	panic("not implemented:" + b.right.Kind().String())
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
	panic("not implemented:" + b.op.String())
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
	}
	panic("not implemented:" + b.op.String())
}

var _ Expr = &operatorUnimplemented{}

type operatorUnimplemented struct{ step }

func (*operatorUnimplemented) Assign(env *Env, value reflect.Value) {
	panic("not implemented")
}
