package internal

import "testing"

func TestStack_PushPop(t *testing.T) {
	s := new(stack)
	if _, ok := s.pop(); ok {
		t.Fatal("pop on empty stack should not be ok")
	}
	s.push(stackFrame{})
	s.push(stackFrame{})
	if len(*s) != 2 {
		t.Fatal("stack should have 2 frames")
	}
	if _, ok := s.pop(); !ok {
		t.Fatal("pop on non-empty stack should be ok")
	}
	if len(*s) != 1 {
		t.Fatal("stack should have 1 frame")
	}
	if _, ok := s.pop(); !ok {
		t.Fatal("pop on non-empty stack should be ok")
	}
	if len(*s) != 0 {
		t.Fatal("stack should be empty")
	}
	if _, ok := s.pop(); ok {
		t.Fatal("pop on empty stack should not be ok")
	}
}

func TestStack_Peek(t *testing.T) {
	s := new(stack)
	if _, ok := s.peek(); ok {
		t.Fatal("peek on empty stack should not be ok")
	}
	s.push(stackFrame{})
	if _, ok := s.peek(); !ok {
		t.Fatal("peek on non-empty stack should be ok")
	}
	if len(*s) != 1 {
		t.Fatal("stack should have 1 frame after peek")
	}
}
