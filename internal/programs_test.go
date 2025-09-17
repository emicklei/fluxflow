package internal

import (
	"os"
	"testing"
)

func TestProgramPrint(t *testing.T) {
	out := parseAndRun(t, `package main

func main() {
	print("flux")
	print("flow")
}
`)
	if got, want := out, "fluxflow"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramMulitAssign(t *testing.T) {
	out := parseAndRun(t, `package main

func main() {
	in1, in2 := "flux", "flow"
	print(in1, in2)
}
`)
	if got, want := out, "fluxflow"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramIfElse(t *testing.T) {
	out := parseAndRun(t, `package main

func main() {
	if 1 == 2 {
		print("unreachable")
	} else {
		print("fluxflow")
	}
}
`)
	if got, want := out, "fluxflow"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramTrueFalse(t *testing.T) {
	out := parseAndRun(t, `package main

func main() {
	print(true, false)
}
`)
	if got, want := out, "truefalse"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramFunc(t *testing.T) {
	out := parseAndRun(t, `package main

func plus(a int, b int) int {
	return a + b
}

func main() {
	result := plus(2, 3)
	print(result)
}`)
	if got, want := out, "5"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramFor(t *testing.T) {
	out := parseAndRun(t, `package main

func main() {
	for i := 0; i < 10; i++ {
		print(i)
	}
}`)
	if got, want := out, "12345678910"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramGeneric(t *testing.T) {
	t.Skip()
	parseAndRun(t, `package main

import "fmt"

func Generic[T any](arg T) (*T, error) { return &arg, nil }

func main() {
	h, _ := Generic("hello")
	fmt.Println(*h)
}

/**
func Generic_string(arg string) (*string, error) { return &arg, nil }

func main() {
	h, _ := Generic_string("hello")
	fmt.Println(*h)
}
**/
`)
}
func TestProgramTypeAssert(t *testing.T) {
	out := parseAndRun(t, `package main

func main() {
	a := int32(1) + 2
	print(a)
}`)
	if got, want := out, "3"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramNil(t *testing.T) {
	t.Skip()
	os.Setenv("STEPS", "1")
	parseAndRun(t, `package main

func main() {
	var s *string = nil
	print(s)
}
`)
}

func TestProgramSlice(t *testing.T) {
	t.Skip()
	os.Setenv("STEPS", "1")
	out := parseAndRun(t, `package main

func main() {
	print([]int{1, 2})
}
`)
	if got, want := out, "[1 2]"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramAppend(t *testing.T) {
	t.Skip()
	os.Setenv("STEPS", "1")
	parseAndRun(t, `package main

func main() {
	list := []int{}
	list = append(list, 4)
	print(list)
}
`)
}
