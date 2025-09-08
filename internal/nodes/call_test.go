package nodes

import (
	"fmt"
	"reflect"
	"testing"
)

func Generic[T any](arg T) (*T, error) { return &arg, nil }

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

func TestCallGenericByReflect(t *testing.T) {
	stringGeneric := Generic[string]
	rv := reflect.ValueOf(stringGeneric)
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

// TestInstantiateGenericByReflect explains that instantiating a generic function
// like `Generic[string]` cannot be done using the standard `reflect` package.
func TestInstantiateGenericByReflect(t *testing.T) {
	// The following is not possible with Go's reflection.
	// You cannot take a generic function definition and supply a type to it at runtime.
	// genericFunc := reflect.ValueOf(Generic)
	// stringType := reflect.TypeOf("")
	// stringGenericFunc := genericFunc.WithType(stringType) // This method does not exist

	// Instantiation of a generic function must be done at compile-time.
	stringGeneric := Generic[string]

	// After compile-time instantiation, you can use reflection on the resulting function.
	rv := reflect.ValueOf(stringGeneric)
	if rv.Kind() != reflect.Func {
		t.Fatalf("expected a function, got %v", rv.Kind())
	}
	t.Logf("type of instantiated generic function: %v", rv.Type())
}
