package problem71

import (
	"strings"
)

const (
	pathSeperator = "/"
	parentDir = ".."
	currentDir = "."
)

func simplifyPath(path string) string {
	paths := strings.Split(path, pathSeperator)
	stack := InitStack[string](len(paths))

	for _, p := range paths {
		if len(p) == 0 || p == currentDir {
			continue
		}
		if p == parentDir {
			stack.Pop()
		} else {
			stack.Push(p)
		}
	}
	
	stackLen := stack.Len()
	s := make([]string, stackLen)
	for i := 0; i < stackLen; i++ {
		s[stackLen-1-i] = stack.Pop()
	}

	return pathSeperator + strings.Join(s, pathSeperator)
}

type stack[T any] struct {
	top int
	container []T
}

func InitStack[T any](cap int) *stack[T] {
	s := &stack[T]{
		container: make([]T, cap),
		top: -1,
	}
	return s
}

func(s *stack[T]) Push(val T) {
	s.top++
	s.container[s.top] = val
}

func (s *stack[T]) Len() int {
	return s.top+1
}

func(s *stack[T]) Pop() (v T) {
	if s.top < 0 {
		return
	}
	v = s.container[s.top]
	s.top--
	return v
}

func(s *stack[T]) Peek() T {
	return s.container[s.top]
}
