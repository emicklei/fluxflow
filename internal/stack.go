package internal

import "reflect"

type stackFrame struct {
	returnStep   Step
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

func (s stack) peek() (stackFrame, bool) {
	if len(s) == 0 {
		return stackFrame{}, false
	}
	return s[len(s)-1], true
}

// pre: stack not empty
func (s stack) top() stackFrame {
	return s[len(s)-1]
}
