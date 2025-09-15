package internal

import "reflect"

type stackFrame struct {
	returnStep   *step
	env          *Env
	localVars    map[string]Var
	funcArgs     []reflect.Value
	returnValues []reflect.Value
}

type stack []stackFrame

func (s *stack) push(f stackFrame) {
	*s = append(*s, f)
}

// pre: stack not empty
func (s *stack) pop() stackFrame {
	f := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return f
}

// pre: stack not empty
func (s stack) top() stackFrame {
	return s[len(s)-1]
}
