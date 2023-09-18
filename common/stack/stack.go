package stack

type Stack[T any] struct {
	container []T
	top int
}

func InitFromSlice[T any](c []T) *Stack[T] {
	return &Stack[T]{
		container: c,
		top: len(c),
	}
}

func Init[T any](cap int) *Stack[T] {
	return &Stack[T]{
		container: make([]T, cap),
		top: -1,
	}
}

func(s *Stack[T]) Push(val T) {
	s.top++
	if s.top >= cap(s.container) {
		s.expand()
	}
	s.container[s.top] = val
}

func (s *Stack[T]) expand() {
	newContainer := make([]T, cap(s.container) << 1)
	copy(newContainer, s.container)

	s.container = nil
	s.container = newContainer
}

func (s *Stack[T]) shrink() {
	halfSize := cap(s.container) >> 1
	newContainer := make([]T, halfSize)
	copy(newContainer,  s.container[:s.top + 1])

	s.container = nil
	s.container = newContainer
}

func (s *Stack[T]) Len() int {
	return s.top+1
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == -1
}

func(s *Stack[T]) Pop() (v T) {
	if s.top < 0 {
		return
	}
	v = s.container[s.top]
	s.top--
	if s.top < (cap(s.container) >> 2) {
		s.shrink()
	}
	return v
}

func(s *Stack[T]) Peek() T {
	return s.container[s.top]
}

