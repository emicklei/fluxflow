package internal

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

type UnaryExpr struct {
	*ast.UnaryExpr
	X Expr
}

func (u UnaryExpr) String() string {
	return fmt.Sprintf("UnaryExpr(%s %s)", u.Op, u.X)
}

func (u UnaryExpr) Eval(vm *VM) {
	v := vm.ReturnsEval(u.X)
	switch v.Kind() {
	case reflect.Int:
		switch u.Op {
		case token.SUB:
			vm.Returns(reflect.ValueOf(int(-v.Int())))
		case token.AND:
			actual := v.Interface()
			vm.Returns(reflect.ValueOf(&actual))
		default:
			panic("missing unary operation on int:" + u.Op.String())
		}
	case reflect.Int64:
		switch u.Op {
		case token.SUB:
			vm.Returns(reflect.ValueOf(-v.Int()))
		default:
			panic("missing unary operation on int64:" + u.Op.String())
		}
	case reflect.Float64:
		switch u.Op {
		case token.SUB:
			vm.Returns(reflect.ValueOf(-v.Float()))
		case token.ADD:
			vm.Returns(reflect.ValueOf(v.Float()))
		default:
			panic("missing unary operation on float64:" + u.Op.String())
		}
	case reflect.Struct:
		switch u.Op {
		case token.AND:
			actual := v.Interface()
			vm.Returns(reflect.ValueOf(&actual))
		default:
			panic("missing unary operation on struct:" + u.Op.String())
		}
	default:
		panic("not implemented: UnaryExpr.Eval:" + v.Kind().String())
	}
}
