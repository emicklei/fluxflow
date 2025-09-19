package internal

import (
	"go/ast"
	"testing"
)

func TestZeroValueOfType(t *testing.T) {
	env := newEnv()
	i := Ident{Ident: &ast.Ident{Name: "string"}}
	v := i.ZeroValue(env)
	if v.Interface() != "" {
		t.Fail()
	}
}
