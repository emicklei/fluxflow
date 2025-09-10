package internal

import (
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
)

type BinaryExpr struct {
	X Operand
	Y Operand
	*ast.BinaryExpr
}

func (s *BinaryExpr) Eval() reflect.Value {
	switch s.Op {
	case token.ADD:
		return s.X.ADD(s.Y)
	case token.SUB:
		return s.X.SUB(s.Y)
	}
	panic("not implemented")
}

type INT struct {
	value string
}

func (i INT) ADD(right Operand) reflect.Value {
	return right.ADD_INT(i)
}
func (i INT) ADD_INT(left INT) reflect.Value {
	li, _ := strconv.Atoi(left.value)
	ri, _ := strconv.Atoi(i.value)
	return reflect.ValueOf(li + ri)
}
func (i INT) SUB(right Operand) reflect.Value {
	return right.SUB_INT(i)
}
func (i INT) SUB_INT(left INT) reflect.Value {
	li, _ := strconv.Atoi(left.value)
	ri, _ := strconv.Atoi(i.value)
	return reflect.ValueOf(li - ri)
}

type Operand interface {
	ADD(right Operand) reflect.Value
	ADD_INT(INT) reflect.Value
	SUB(right Operand) reflect.Value
	SUB_INT(INT) reflect.Value
}
