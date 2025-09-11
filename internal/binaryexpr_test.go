package internal

import (
	"go/ast"
	"go/token"
	"reflect"
	"testing"
)

// 1 + 2
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

// 5 - 3
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

// "Hello" + "World!"
func TestBinaryStringAddition(t *testing.T) {
	left := Literal{BasicLit: &ast.BasicLit{Kind: token.STRING, Value: "Hello, "}}
	right := Literal{BasicLit: &ast.BasicLit{Kind: token.STRING, Value: "World!"}}
	expr := BinaryExpr{
		X:          left.Operand(),
		Y:          right.Operand(),
		BinaryExpr: &ast.BinaryExpr{Op: token.ADD},
	}
	result := expr.Eval()
	if result.Kind() != reflect.String {
		t.Fatalf("expected string result, got %v", result.Kind())
	}
	if result.String() != "Hello, World!" {
		t.Fatalf(`expected "Hello, World!", got %s`, result.String())
	}
}
func TestBinaryAddFloatToInteger(t *testing.T) {
	left := Literal{BasicLit: &ast.BasicLit{Kind: token.FLOAT, Value: "3.14"}}
	right := Literal{BasicLit: &ast.BasicLit{Kind: token.INT, Value: "42"}}
	expr := BinaryExpr{
		X:          left.Operand(),
		Y:          right.Operand(),
		BinaryExpr: &ast.BinaryExpr{Op: token.ADD},
	}
	result := expr.Eval()
	if result.Kind() != reflect.Float64 {
		t.Fatalf("expected float64 result, got %v", result.Kind())
	}
	if result.Float() != 45.14 {
		t.Fatalf("expected 45.14, got %f", result.Float())
	}
}
