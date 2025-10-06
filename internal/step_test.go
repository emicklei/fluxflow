package internal

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestStep(t *testing.T) {
	t.Run("Next", func(t *testing.T) {
		s1 := &step{}
		s2 := &step{}
		s1.Next(s2)
		if s1.next != s2 {
			t.Error("s1.next should be s2")
		}
		if s2.prev != s1 {
			t.Error("s2.prev should be s1")
		}
		// test idempotency
		s1.Next(s2)
		if s1.next != s2 {
			t.Error("s1.next should be s2")
		}
		if s2.prev != s1 {
			t.Error("s2.prev should be s1")
		}
	})

	t.Run("Prev", func(t *testing.T) {
		s1 := &step{}
		s2 := &step{}
		s2.Prev(s1)
		if s2.prev != s1 {
			t.Error("s2.prev should be s1")
		}
		if s1.next != s2 {
			t.Error("s1.next should be s2")
		}
		// test idempotency
		s2.Prev(s1)
		if s2.prev != s1 {
			t.Error("s2.prev should be s1")
		}
		if s1.next != s2 {
			t.Error("s1.next should be s2")
		}
	})
}

func TestStepByStep(t *testing.T) {
	left := BasicLit{BasicLit: &ast.BasicLit{Kind: token.STRING, Value: "Hello, "}}
	right := BasicLit{BasicLit: &ast.BasicLit{Kind: token.STRING, Value: "World!"}}
	expr := BinaryExpr{
		X:          left,
		Y:          right,
		BinaryExpr: &ast.BinaryExpr{Op: token.ADD},
	}
	leftStep := &step{Evaluable: left}
	rightStep := &step{Evaluable: right}
	rightStep.Prev(leftStep)
	binExprStep := &step{Evaluable: expr}
	binExprStep.Prev(rightStep)

	vm := newVM(newEnvironment(nil))
	here := leftStep
	for here != nil {
		t.Log(here)
		here.Eval(vm)
		here = here.next
	}
	t.Log("result:", vm.callStack.top().pop().Interface())
}
