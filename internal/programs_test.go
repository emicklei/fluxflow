package internal

import (
	"testing"
)

func TestProgramPrint(t *testing.T) {
	loadAndRun(t, "../programs/test_print")
}

func TestProgramMulitAssign(t *testing.T) {
	loadAndRun(t, "../programs/test_multiassign")
}

func TestProgramGeneric(t *testing.T) {
	loadAndRun(t, "../programs/test_generic")
}
func TestProgramTypeAssert(t *testing.T) {
	loadAndRun(t, "../programs/test_typeassert")
}

func TestProgramIf(t *testing.T) {
	loadAndRun(t, "../programs/test_if")
}

func TestProgramFunc(t *testing.T) {
	loadAndRun(t, "../programs/test_func")
}

func TestProgramFor(t *testing.T) {
	out := loadAndRun(t, "../programs/test_for")
	t.Log(out)
}
