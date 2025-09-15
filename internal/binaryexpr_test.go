package internal

import (
	"fmt"
	"go/ast"
	"go/token"
	"math"
	"reflect"
	"strings"
	"testing"
)

func TestBinaryExprValue_Eval(t *testing.T) {
	t.Run("int op int", func(t *testing.T) {
		cases := []struct {
			op       token.Token
			left     string
			right    string
			expected any
		}{
			{token.ADD, "1", "2", int64(3)},
			{token.SUB, "5", "3", int64(2)},
			{token.MUL, "2", "3", int64(6)},
			{token.QUO, "6", "3", int64(2)},
			{token.REM, "7", "3", int64(1)},
			{token.EQL, "3", "3", true},
			{token.EQL, "3", "4", false},
			{token.NEQ, "3", "4", true},
			{token.NEQ, "3", "3", false},
			{token.LSS, "3", "4", true},
			{token.LSS, "3", "3", false},
			{token.LEQ, "3", "3", true},
			{token.LEQ, "3", "4", true},
			{token.LEQ, "4", "3", false},
			{token.GTR, "4", "3", true},
			{token.GTR, "3", "3", false},
			{token.GEQ, "4", "4", true},
			{token.GEQ, "4", "3", true},
			{token.GEQ, "3", "4", false},
		}
		for _, tt := range cases {
			t.Run(fmt.Sprintf("int %s int", tt.op), func(t *testing.T) {
				left := BasicLit{BasicLit: &ast.BasicLit{Kind: token.INT, Value: tt.left}}
				right := BasicLit{BasicLit: &ast.BasicLit{Kind: token.INT, Value: tt.right}}
				expr := BinaryExpr{
					X:          left,
					Y:          right,
					BinaryExpr: &ast.BinaryExpr{Op: tt.op},
				}
				vm := newVM()
				result := vm.ReturnsEval(expr)
				switch expected := tt.expected.(type) {
				case int64:
					if result.Int() != expected {
						t.Errorf("expected %d, got %d", expected, result.Int())
					}
				case bool:
					if result.Bool() != expected {
						t.Errorf("expected %v, got %v", expected, result.Bool())
					}
				}
			})
		}
	})

	t.Run("float op float", func(t *testing.T) {
		cases := []struct {
			op       token.Token
			left     string
			right    string
			expected float64
		}{
			{token.ADD, "1.1", "2.2", 3.3},
			{token.SUB, "5.5", "3.3", 2.2},
			{token.MUL, "2.2", "3.3", 7.26},
			{token.QUO, "6.6", "3.3", 2.0},
		}
		for _, tt := range cases {
			t.Run(fmt.Sprintf("float %s float", tt.op), func(t *testing.T) {
				left := BasicLit{BasicLit: &ast.BasicLit{Kind: token.FLOAT, Value: tt.left}}
				right := BasicLit{BasicLit: &ast.BasicLit{Kind: token.FLOAT, Value: tt.right}}
				expr := BinaryExpr{
					X:          left,
					Y:          right,
					BinaryExpr: &ast.BinaryExpr{Op: tt.op},
				}
				vm := newVM()
				expr.Eval(vm)
				result := vm.ReturnsEval(expr)
				if math.Abs(result.Float()-tt.expected) > 1e-9 {
					t.Errorf("expected %f, got %f", tt.expected, result.Float())
				}
			})
		}
	})

	t.Run("int op float", func(t *testing.T) {
		left := BasicLit{BasicLit: &ast.BasicLit{Kind: token.INT, Value: "42"}}
		right := BasicLit{BasicLit: &ast.BasicLit{Kind: token.FLOAT, Value: "3.14"}}
		expr := BinaryExpr{
			X:          left,
			Y:          right,
			BinaryExpr: &ast.BinaryExpr{Op: token.ADD},
		}
		vm := newVM()
		expr.Eval(vm)
		result := vm.ReturnsEval(expr)
		if result.Kind() != reflect.Float64 {
			t.Fatalf("expected float64 result, got %v", result.Kind())
		}
		if result.Float() != 45.14 {
			t.Fatalf("expected 45.14, got %f", result.Float())
		}
	})

	t.Run("float op int", func(t *testing.T) {
		left := BasicLit{BasicLit: &ast.BasicLit{Kind: token.FLOAT, Value: "3.14"}}
		right := BasicLit{BasicLit: &ast.BasicLit{Kind: token.INT, Value: "42"}}
		expr := BinaryExpr{
			X:          left,
			Y:          right,
			BinaryExpr: &ast.BinaryExpr{Op: token.ADD},
		}
		vm := newVM()
		expr.Eval(vm)
		result := vm.ReturnsEval(expr)
		if result.Kind() != reflect.Float64 {
			t.Fatalf("expected float64 result, got %v", result.Kind())
		}
		if result.Float() != 45.14 {
			t.Fatalf("expected 45.14, got %f", result.Float())
		}
	})

	t.Run("string op string", func(t *testing.T) {
		left := BasicLit{BasicLit: &ast.BasicLit{Kind: token.STRING, Value: "Hello, "}}
		right := BasicLit{BasicLit: &ast.BasicLit{Kind: token.STRING, Value: "World!"}}
		expr := BinaryExpr{
			X:          left,
			Y:          right,
			BinaryExpr: &ast.BinaryExpr{Op: token.ADD},
		}
		vm := newVM()
		expr.Eval(newVM())
		result := vm.ReturnsEval(expr)
		if result.Kind() != reflect.String {
			t.Fatalf("expected string result, got %v", result.Kind())
		}
		if result.String() != "Hello, World!" {
			t.Fatalf(`expected "Hello, World!", got %s`, result.String())
		}
	})

	t.Run("panics", func(t *testing.T) {
		cases := []struct {
			name          string
			op            token.Token
			left          Expr
			right         Expr
			panicContains string
		}{
			{
				name:          "invalid left kind",
				op:            token.ADD,
				left:          &mockExpr{v: reflect.ValueOf(true)},
				right:         &mockExpr{v: reflect.ValueOf(1)},
				panicContains: "not implemented:",
			},
			{
				name:          "string sub",
				op:            token.SUB,
				left:          &mockExpr{v: reflect.ValueOf("a")},
				right:         &mockExpr{v: reflect.ValueOf("b")},
				panicContains: "not implemented:",
			},
			{
				name:          "string add int",
				op:            token.ADD,
				left:          &mockExpr{v: reflect.ValueOf("a")},
				right:         &mockExpr{v: reflect.ValueOf(1)},
				panicContains: "not implemented:",
			},
			{
				name:          "int add string",
				op:            token.ADD,
				left:          &mockExpr{v: reflect.ValueOf(1)},
				right:         &mockExpr{v: reflect.ValueOf("a")},
				panicContains: "not implemented:",
			},
			{
				name:          "float add string",
				op:            token.ADD,
				left:          &mockExpr{v: reflect.ValueOf(1.0)},
				right:         &mockExpr{v: reflect.ValueOf("a")},
				panicContains: "not implemented:",
			},
			{
				name:          "float rem float",
				op:            token.REM,
				left:          &mockExpr{v: reflect.ValueOf(1.0)},
				right:         &mockExpr{v: reflect.ValueOf(2.0)},
				panicContains: "not implemented:",
			},
			{
				name:          "int and int",
				op:            token.AND,
				left:          &mockExpr{v: reflect.ValueOf(1)},
				right:         &mockExpr{v: reflect.ValueOf(2)},
				panicContains: "not implemented:",
			},
		}
		for _, tt := range cases {
			t.Run(tt.name, func(t *testing.T) {
				defer func() {
					r := recover()
					if r == nil {
						t.Fatal("did not panic")
					}
					msg := fmt.Sprintf("%v", r)
					if !strings.Contains(msg, tt.panicContains) {
						t.Errorf("expected panic message to contain %q, got %q", tt.panicContains, msg)
					}
				}()
				expr := BinaryExpr{
					X:          tt.left,
					Y:          tt.right,
					BinaryExpr: &ast.BinaryExpr{Op: tt.op},
				}
				expr.Eval(newVM())
			})
		}
	})
}

type mockExpr struct {
	v reflect.Value
}

func (m *mockExpr) Eval(vm *VM) {
}
func (m *mockExpr) String() string { return "mock" }
