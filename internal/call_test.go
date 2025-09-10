package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCallByReflect(t *testing.T) {
	rv := reflect.ValueOf(fmt.Println)
	rv1 := reflect.ValueOf("fluxflow")
	rv2 := reflect.ValueOf("in")
	rv3 := reflect.ValueOf("out")
	args := []reflect.Value{rv1, rv2, rv3}
	results := rv.Call(args)
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
}

func Generic[T any](arg T) (*T, error) { return &arg, nil }

// instantiations of Generic
func Generic_string(arg string) (*string, error) { return &arg, nil }

func TestCallGenericByReflect(t *testing.T) {
	rv := reflect.ValueOf(Generic_string)
	arg := "hello"
	rvArg := reflect.ValueOf(arg)
	results := rv.Call([]reflect.Value{rvArg})
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
	if results[0].IsNil() {
		t.Fatal("result pointer should not be nil")
	}
	sPtr := results[0].Interface().(*string)
	if *sPtr != arg {
		t.Errorf("expected result to be %q, got %q", arg, *sPtr)
	}
	if !results[1].IsNil() {
		t.Errorf("expected error to be nil, got %v", results[1].Interface())
	}
}

// a := "a"
// fmt.Println(a)
func TestCallWithStackframe(t *testing.T) {
	// f := stackFrame{}
	// f.funcArgs = []reflect.Value{reflect.ValueOf("a")}

}
