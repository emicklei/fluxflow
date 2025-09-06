package nodes

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
