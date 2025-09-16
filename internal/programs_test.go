package internal

import (
	"os"
	"testing"
)

func TestProgramPrint(t *testing.T) {
	out := loadAndRun(t, "../programs/test_print")
	if got, want := out, "fluxflow"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramMulitAssign(t *testing.T) {
	out := loadAndRun(t, "../programs/test_multiassign")
	if got, want := out, "fluxflow"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramIfElse(t *testing.T) {
	out := loadAndRun(t, "../programs/test_if_else")
	if got, want := out, "fluxflow"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramTrueFalse(t *testing.T) {
	out := loadAndRun(t, "../programs/test_true_false")
	if got, want := out, "truefalse"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramFunc(t *testing.T) {
	out := loadAndRun(t, "../programs/test_func")
	if got, want := out, "5"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramFor(t *testing.T) {
	out := loadAndRun(t, "../programs/test_for")
	if got, want := out, "12345678910"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramGeneric(t *testing.T) {
	t.Skip()
	loadAndRun(t, "../programs/test_generic")
}
func TestProgramTypeAssert(t *testing.T) {
	out := loadAndRun(t, "../programs/test_builtin_type_convert")
	if got, want := out, "3"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramNil(t *testing.T) {
	t.Skip()
	os.Setenv("STEPS", "1")
	loadAndRun(t, "../programs/test_nil")
}
