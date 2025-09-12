package internal

import (
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
)

type BinaryExpr struct {
	step
	X Expr
	Y Expr
	*ast.BinaryExpr
}

func (s BinaryExpr) Eval(env *Env) reflect.Value {
	switch s.Op {
	case token.ADD:
		return s.X.ADD(s.Y)
	case token.SUB:
		return s.X.SUB(s.Y)
	}
	panic("not implemented")
}

type INT struct {
	operatorUnimplemented
	value string
}

func (i INT) ADD(right Expr) reflect.Value {
	return right.ADD_INT(i)
}
func (i INT) ADD_INT(left INT) reflect.Value {
	li, _ := strconv.Atoi(left.value)
	ri, _ := strconv.Atoi(i.value)
	return reflect.ValueOf(li + ri)
}
func (i INT) SUB(right Expr) reflect.Value {
	return right.SUB_INT(i)
}
func (i INT) SUB_INT(left INT) reflect.Value {
	li, _ := strconv.Atoi(left.value)
	ri, _ := strconv.Atoi(i.value)
	return reflect.ValueOf(li - ri)
}
func (i INT) ADD_FLOAT(left FLOAT) reflect.Value {
	li, _ := strconv.Atoi(i.value)
	ri, _ := strconv.ParseFloat(left.value, 64)
	return reflect.ValueOf(float64(li) + ri)
}

type Expr interface {
	Eval(env *Env) reflect.Value

	ADD(right Expr) reflect.Value
	ADD_INT(INT) reflect.Value

	SUB(right Expr) reflect.Value
	SUB_INT(INT) reflect.Value

	ADD_STRING(STRING) reflect.Value

	ADD_FLOAT(FLOAT) reflect.Value
	SUB_FLOAT(FLOAT) reflect.Value

	Assign(env *Env, val reflect.Value)
}

var _ Expr = operatorUnimplemented{}

type operatorUnimplemented struct{ step }

func (o operatorUnimplemented) ADD(right Expr) reflect.Value {
	panic("not implemented")
}
func (o operatorUnimplemented) ADD_INT(left INT) reflect.Value {
	panic("not implemented")
}
func (o operatorUnimplemented) ADD_STRING(left STRING) reflect.Value {
	panic("not implemented")
}
func (o operatorUnimplemented) SUB(right Expr) reflect.Value {
	panic("not implemented")
}
func (o operatorUnimplemented) SUB_INT(left INT) reflect.Value {
	panic("not implemented")
}
func (o operatorUnimplemented) ADD_FLOAT(left FLOAT) reflect.Value {
	panic("not implemented")
}
func (o operatorUnimplemented) SUB_FLOAT(left FLOAT) reflect.Value {
	panic("not implemented")
}
func (o operatorUnimplemented) Assign(env *Env, val reflect.Value) {
	panic("not implemented")
}

type STRING struct {
	operatorUnimplemented
	value string
}

func (s STRING) ADD(right Expr) reflect.Value {
	return right.ADD_STRING(s)
}
func (s STRING) ADD_STRING(left STRING) reflect.Value {
	return reflect.ValueOf(left.value + s.value)
}

type FLOAT struct {
	operatorUnimplemented
	value string
}

func (f FLOAT) ADD(right Expr) reflect.Value {
	return right.ADD_FLOAT(f)
}
func (f FLOAT) ADD_FLOAT(left FLOAT) reflect.Value {
	li, _ := strconv.ParseFloat(left.value, 64)
	ri, _ := strconv.ParseFloat(f.value, 64)
	return reflect.ValueOf(li + ri)
}

func (f FLOAT) SUB(right Expr) reflect.Value {
	return right.SUB_FLOAT(f)
}
func (f FLOAT) SUB_FLOAT(left FLOAT) reflect.Value {
	li, _ := strconv.ParseFloat(left.value, 64)
	ri, _ := strconv.ParseFloat(f.value, 64)
	return reflect.ValueOf(li - ri)
}
func (f FLOAT) ADD_INT(left INT) reflect.Value {
	li, _ := strconv.Atoi(left.value)
	ri, _ := strconv.ParseFloat(f.value, 64)
	return reflect.ValueOf(float64(li) + ri)
}
