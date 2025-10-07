package internal

import (
	"fmt"
	"strings"
	"testing"
)

func TestProgramTypeConvert(t *testing.T) {
	t.Skip()
	tests := []struct {
		typeName string
	}{
		{"int8"},
		{"int16"},
		{"int32"},
		{"int64"},
		{"int"},
	}
	for _, tt := range tests {
		t.Run(tt.typeName, func(t *testing.T) {
			out := parseAndRun(t, fmt.Sprintf(`
package main

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

func TestProgramTypeUnsignedConvert(t *testing.T) {
	// t.Skip()
	tests := []struct {
		typeName string
	}{
		{"uint8"},
		{"uint16"},
		{"uint32"},
		{"uint64"},
		{"uint"},
	}
	for _, tt := range tests {
		t.Run(tt.typeName, func(t *testing.T) {
			out := parseAndRun(t, fmt.Sprintf(`
package main

func main() {
	a := %s(1) + %s(2)
	print(a)
}`, tt.typeName, tt.typeName))
			if got, want := out, "3"; got != want {
				t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
			}
		})
	}
}

func TestAssignmentOperators(t *testing.T) {
	tests := []struct {
		op   string
		want string
	}{
		{"+=", "3"},
		{"-=", "-1"},
		{"*=", "2"},
		{"/=", "0"},
		{"%=", "1"},
		{"&=", "0"},
		{"|=", "3"},
		{"^=", "3"},
		{"<<=", "4"},
		{">>=", "0"},
		{"&^=", "1"},
	}
	for _, tt := range tests {
		t.Run(tt.op, func(t *testing.T) {
			out := parseAndRun(t, fmt.Sprintf(`
package main

func main() {
	a := 1
	a %s 2
	print(a)
}`, tt.op))
			if got, want := out, tt.want; got != want {
				t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
			}
		})
	}
}

func TestAllPrograms(t *testing.T) {
	tests := []struct {
		name    string
		source  string
		want    string
		skip    bool
		debug   bool
		special func(string) bool // special check for output
	}{
		{
			name: "print",
			source: `
package main

func main() {
	print("flux")
	print("flow")
}`,
			want: "fluxflow",
		},
		{
			name: "multi-assign",
			source: `
package main

func main() {
	in1, in2 := "flux", "flow"
	print(in1, in2)
}`,
			want: "fluxflow",
		},
		{
			name: "if-else",
			source: `
package main

func main() {
	if 1 == 2 {
		print("unreachable")
	} else {
		print("fluxflow")
	}
}`,
			want: "fluxflow",
		},
		{
			name: "true-false",
			source: `
package main

func main() {
	print(true, false)
}`,
			want: "truefalse",
		},
		{
			name: "rune",
			source: `
package main

func main() {
	print('e')
}`,
			want: "'e'",
		},
		{
			name: "numbers",
			source: `
package main

func main() {
	print(-1,+3.14,0.1e10)
}`,
			want: "-13.141e+09",
		},
		{
			name: "func",
			source: `
package main

func plus(a int, b int) int {
	return a + b
}
func main() {
	result := plus(2, 3)
	print(result)
}`,
			want: "5",
		},
		{
			name: "func-multi-return",
			source: `
package main

func ab(a int, b int) (int,int) {
	return a,b
}
func main() {
	a,b := ab(2, 3)
	print(a,b)
}`,
			want: "23",
		},
		{
			name: "for",
			source: `
package main

func main() {
	for i := 0; i < 10; i++ {
		print(i)
	}
	for i := 9; i > 0; i-- {
		print(i)
	}		
}`,
			want: "0123456789987654321",
		},
		{
			name: "for-scope",
			source: `
package main

func main() {
	j := 1
	for i := 0; i < 3; i++ {
		j = i
		print(i)
	}
	print(j)
}`,
			want: "0122",
		},
		{
			name: "generic",
			skip: true,
			source: `
package main

import "fmt"
func Generic[T any](arg T) (*T, error) { return &arg, nil }
func main() {
	h, _ := Generic("hello")
	fmt.Println(*h)
}`,
		},
		{
			name: "declare",
			source: `
package main

func main() {
	var s string
	print(s)
}`,
			want: "",
		},
		{
			name: "const",
			source: `
package main

const (
	C = A+1
	A = 0
	B = 1	
)
func main() {	
	print(A,B,C)
}`,
			want: "011",
		},
		{
			name: "var",
			source: `
package main

var (
	a = 1
	s string
	b bool
)
func main() {	
	print(a,s,b)
}`,
			want: "1false",
		},
		{
			name: "const-scope",
			skip: true,
			source: `
package main

var b = a
const a = 1
func main() {
	var b = a
	const a = 2
	print(a, b)
}`,
			want: "21",
		},
		{
			name: "declare-and-init",
			source: `
package main

func main() {
	var s string = "fluxflow"
	print(s)
}`,
			want: "fluxflow",
		},
		{
			name: "slice",
			source: `
package main

func main() {
	print([]int{1, 2})
}`,
			want: "[1 2]",
		},
		{
			name: "array",
			source: `
package main

func main() {
	print([2]string{"A", "B"})
}`,
			want: "[A B]",
		},
		{
			name: "slice-append-and-index",
			source: `
package main

func main() {
	list := []int{}
	list = append(list, 1, 2)
	print(list[0], list[1])
}`,
			want: "12",
		},
		{
			name: "append",
			source: `
package main

func main() {
	list := []int{}
	print(list)
	list = append(list, 4, 5)
	print(list)
}`,
			want: "[][4 5]",
		},
		{
			name: "time-constant",
			source: `
package main

import "time"
func main() {
	r := time.RFC1123
	print(r)
}`,
			want: "Mon, 02 Jan 2006 15:04:05 MST",
		},
		{
			name: "time-alias-constant",
			source: `
package main

import t "time"
func main() {
	r := t.RFC1123
	print(r)
}`,
			want: "Mon, 02 Jan 2006 15:04:05 MST",
		},
		{
			name: "floats",
			source: `
package main

func main() {
	f32, f64 := float32(3.14), 3.14
	print(f32," ",f64)
}`,
			want: "3.14 3.14",
		},
		{
			name: "new-type",
			source: `
package main

type Airplane struct {
	Kind string
}
func main() {
	heli := Airplane{Kind:"helicopter"}
	print(heli.Kind)
}`,
			want: "helicopter",
		},
		{
			name: "pointer-to-type",
			skip: true,
			source: `
package main

type Airplane struct {
	Kind string
}
func main() {
	heli := &Airplane{Kind:"helicopter"}
	print(heli.Kind)
}`,
			want: "helicopter",
		},
		{
			name: "address-of-int",
			source: `
package main

func main() {
	i := 42
	print(&i)
}`,
			special: func(out string) bool { return strings.HasPrefix(out, "0x") },
		},
		{
			name: "range-of-strings",
			source: `
package main

func main() {
	strings := []string{"hello", "world"}
	for i,s := range strings {
		print(i,s)
	}
}`,
			want: "0hello1world",
		},
		{
			name: "init",
			source: `
package main

func init() {
	print("0")
}
func init() {
	print("1")
}
func main() {}`,
			want: "01",
		},
		{
			name: "method",
			skip: true,
			source: `
package main

func (_ Airplane) S() string { return "airplane" }
type Airplane struct {}
func main() {
	print(Airplane{}.S())
}`,
			want: "airplane",
		},
		{
			name: "goto",
			skip: true,
			source: `
package main

func main() {
	s := 1
one:
	print(s)
	s++
	if s == 4 {
		return
	} else {
		goto two
	}
two:
	print(s)
	s++
	goto one
}
`,
			want: "aaa",
		},
		{
			name: "map",
			source: `
package main

func main() {
	m := map[string]int{}
	m["a"] = 1
	m["b"] = 2
	print(m["a"] + m["b"])
}`,
			want: "3",
		},
		{
			name: "switch",
			source: `
package main

func main() {
	var a int
	switch a = 1; a {
	case 1:
		print(a)
	}
	switch a {
	case 2:
	default:
		print(2)
	}
}`,
			want: "12",
		},
		{
			name:  "switch-on-bool",
			debug: !true,
			source: `
package main

func main() {
	var a int = 1
	switch {
	case a == 1:
		print(a)
	}
}`,
			want: "1",
		},
		{
			name:  "function-literal",
			skip:  !true,
			debug: !true,
			source: `
package main

func main() {
	f := func(a int) int { return a } 
	print(f(1))
}`,
			want: "1",
		},
		{
			name:  "type-alias",
			skip:  !true,
			debug: !true,
			source: `
package main

type MyInt = int

func main() {
	var a MyInt = 1
	print(a)
}`,
			want: "1",
		},
		{
			name:  "defer",
			skip:  !true,
			debug: !true,
			source: `
package main

func main() {
	defer print(1)
	defer print(2)
}`,
			want: "12",
		},
		{
			name:  "rune",
			skip:  !true,
			debug: !true,
			source: `
package main

func main() {
	r := 'a'
	print(r)
}`,
			want: "'a'",
		},
		{
			name:  "variadic-function",
			skip:  true,
			debug: !true,
			source: `
package main

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	print(sum(1, 2, 3))
}`,
			want: "6",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skip()
			}
			if tt.debug {
				t.Log("debug and trace on for", tt.name)
				// put a breakpoint here to inspect the test case
				trace = true
			}
			out := parseAndRun(t, tt.source)
			if tt.special != nil {
				if !tt.special(out) {
					t.Errorf("special check failed on output: %s", out)
				}
				return
			}
			if got, want := out, tt.want; got != want {
				t.Errorf("got [%v] want [%v]", got, want)
			}
		})
	}
}

func TestIsolatedProgram(t *testing.T) {
	prog := buildProgram(t, `
package main

func main() {
	print("one")
	print("two")
}`)
	// vm := newVM(prog.builder.env)
	// if err := RunProgram(prog, vm); err != nil {
	// 	t.Fatal(err)
	// }
	var here Step = prog.builder.stack[0]
	for here != nil {
		t.Log(here)
		here = here.Next()
	}
}
