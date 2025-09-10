package internal

import (
	"go/ast"
	"go/token"
	"reflect"
	"testing"
)

func TestBinaryIntegerAddition(t *testing.T) {
	left := Literal{BasicLit: &ast.BasicLit{Kind: token.INT, Value: "1"}}
	right := Literal{BasicLit: &ast.BasicLit{Kind: token.INT, Value: "2"}}
	expr := BinaryExpr{
		X:          left.Operand(),
		Y:          right.Operand(),
		BinaryExpr: &ast.BinaryExpr{Op: token.ADD},
	}
	result := expr.Eval()
	if result.Kind() != reflect.Int {
		t.Fatalf("expected int result, got %v", result.Kind())
	}
	if result.Int() != 3 {
		t.Fatalf("expected 3, got %d", result.Int())
	}
}
func TestBinaryIntegerSubtraction(t *testing.T) {
	left := Literal{BasicLit: &ast.BasicLit{Kind: token.INT, Value: "5"}}
	right := Literal{BasicLit: &ast.BasicLit{Kind: token.INT, Value: "3"}}
	expr := BinaryExpr{
		X:          left.Operand(),
		Y:          right.Operand(),
		BinaryExpr: &ast.BinaryExpr{Op: token.SUB},
	}
	result := expr.Eval()
	if result.Kind() != reflect.Int {
		t.Fatalf("expected int result, got %v", result.Kind())
	}
	if result.Int() != 2 {
		t.Fatalf("expected 2, got %d", result.Int())
	}
}
