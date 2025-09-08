package internal

import "reflect"

type stackFrame struct {
	returnNode   node
	localVars    map[string]reflect.Value
	funcArgs     []reflect.Value
	returnValues []reflect.Value
}

type stack []stackFrame

func (s *stack) push(f stackFrame) {
	*s = append(*s, f)
}

func (s *stack) pop() (stackFrame, bool) {
	if len(*s) == 0 {
		return stackFrame{}, false
	}
	f := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return f, true
}

func (s stack) peek() (stackFrame, bool) {
	if len(s) == 0 {
		return stackFrame{}, false
	}
	return s[len(s)-1], true
}
