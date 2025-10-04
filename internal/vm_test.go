package internal

import (
	"reflect"
	"testing"
)

func TestStackFramePushPop(t *testing.T) {
	// Create a new stackFrame with a nil environment.
	frame := stackFrame{env: newEnvironment(nil)}

	// The value to push onto the stack.
	want := reflect.ValueOf(42)

	// Push the value onto the stack.
	frame.push(want)

	// Pop the value from the stack.
	got := frame.pop()

	// Assert that the popped value is the same as the pushed value.
	if got.Interface() != want.Interface() {
		t.Errorf("pop() = %v, want %v", got.Interface(), want.Interface())
	}
}
