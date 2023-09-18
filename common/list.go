package common

type IValue = any

type ListNode[T IValue] struct {
	Val T
	Next *ListNode[T]
}

func NewListNode[T IValue](val T, next *ListNode[T]) *ListNode[T] {
	return &ListNode[T]{
		Val: val,
		Next: next,
	}
}

type List[T IValue] struct {
	Head *ListNode[T]
	Size int
}

func FromSlice[T IValue](slice []T) *List[T] {
	if len(slice) == 0 {
		return &List[T]{
			Head: nil,
			Size: 0,
		}
	}

	head := NewListNode(slice[0], nil)
	cursor := head
	for i := 1; i < len(slice); i++ {
		new := NewListNode(slice[i], nil)
		cursor.Next = new
		cursor = cursor.Next
	}

	return &List[T]{Head: head, Size: len(slice)}
}

func(l *List[T]) ToSlice() []T {
	slice := make([]T, 0, l.Size)
	cursor := l.Head

	for cursor != nil {
		slice = append(slice, cursor.Val)
		cursor = cursor.Next
	}
	return slice
}
