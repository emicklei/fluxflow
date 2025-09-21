package internal

import (
	"fmt"
	"testing"
)

func TestProgramEverything(t *testing.T) {
	t.Skip()
	t.Log(loadAndRun(t, "../programs"))
}

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

func TestProgramRune(t *testing.T) {
	out := parseAndRun(t, `package main

func main() {
	print('e')
}`)
	if got, want := out, "'e'"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramNumbers(t *testing.T) {
	printSteps()
	out := parseAndRun(t, `package main

func main() {
	print(-1,+3.14,0.1e10)
}`)
	if got, want := out, "-13.141e+09"; got != want {
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

func TestProgramFuncMultiReturn(t *testing.T) {
	t.Skip()
	printSteps()
	out := parseAndRun(t, `package main

func ab(a int, b int) (int,int) {
	return a,b
}

func main() {
	a,b := ab(2, 3)
	print(a,b)
}`)
	if got, want := out, "23"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramFor(t *testing.T) {
	out := parseAndRun(t, `package main

func main() {
	for i := 0; i < 10; i++ {
		print(i)
	}
	for i := 9; i > 0; i-- {
		print(i)
	}		
}`)
	if got, want := out, "0123456789987654321"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
func TestProgramForScope(t *testing.T) {
	out := parseAndRun(t, `package main

func main() {
	j := 1
	for i := 0; i < 3; i++ {
		j = i
		print(i)
	}
	print(j)
}`)
	if got, want := out, "0122"; got != want {
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
func TestProgramTypeConvert(t *testing.T) {
	tests := []struct {
		typeName string
	}{
		{"int8"},
		{"int16"},
		{"int32"},
		{"int64"},
	}
	for _, tt := range tests {
		t.Run(tt.typeName, func(t *testing.T) {
			out := parseAndRun(t, fmt.Sprintf(`package main
func main() {
	a := %s(1) + 2
	print(a)
}`, tt.typeName))
			if got, want := out, "3"; got != want {
				t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
			}
		})
	}
}

func TestProgramDeclare(t *testing.T) {
	out := parseAndRun(t, `package main

func main() {
	var s string
	print(s)
}
`)
	if got, want := out, ""; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestProgramDeclareAndInit(t *testing.T) {
	out := parseAndRun(t, `package main

func main() {
	var s string = "fluxflow"
	print(s)
}
`)
	if got, want := out, "fluxflow"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
func TestProgramSlice(t *testing.T) {
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
	out := parseAndRun(t, `package main

func main() {
	list := []int{}
	print(list)
	list = append(list, 4, 5)
	print(list)
}
`)
	if got, want := out, "[][4 5]"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestTimeConstant(t *testing.T) {
	out := parseAndRun(t, `package main

import "time"

func main() {
	r := time.RFC1123
	print(r)
}
`)
	if got, want := out, "Mon, 02 Jan 2006 15:04:05 MST"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
func TestTimeAliasConstant(t *testing.T) {
	out := parseAndRun(t, `package main

import t "time"

func main() {
	r := t.RFC1123
	print(r)
}
`)
	if got, want := out, "Mon, 02 Jan 2006 15:04:05 MST"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestFloats(t *testing.T) {
	out := parseAndRun(t, `package main

func main() {
	f32, f64 := float32(3.14), 3.14
	print(f32," ",f64)
}
`)
	if got, want := out, "3.14 3.14"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
