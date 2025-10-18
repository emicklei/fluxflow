package internal

import (
	"reflect"
	"testing"
)

func TestStackFramePushPop(t *testing.T) {
	r1 := reflect.ValueOf(42)
	f0 := stackFrame{}
	f1 := f0.withOperand(r1)
	if got, want := len(f0.operandStack), 0; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
	if got, want := len(f1.operandStack), 1; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
	f2 := f1.withReturnValue(r1)
	if got, want := len(f1.returnValues), 0; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
	if got, want := len(f2.returnValues), 1; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
