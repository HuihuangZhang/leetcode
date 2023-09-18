package tree

type Node[T comparable] struct {
    Val T
    Left *Node[T]
    Right *Node[T]
}

func ExportToSliceByLevelTraversal[T comparable](root *Node[T], nullVal T) []T {
	queue := []*Node[T]{}
	result := []T{}

	if root == nil {
		return result
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		first := queue[0]
		queue = queue[1:]

		if first == nil {
			result = append(result, nullVal)
			continue
		} else {
			result = append(result, first.Val)
		}

		if first.Left == nil && first.Right == nil {
			continue
		}
		queue = append(queue, first.Left)
		queue = append(queue, first.Right)
	}
	return result
}

func ImportFromSlice[T comparable](s []T, nullVal T) *Node[T] {
	nodeSlice := make([]*Node[T], 0, len(s))
	if len(s) == 0 {
		return nil
	}
	nodeSlice = append(nodeSlice, NewNode(s[0], nil, nil))
	for i := 1; i < len(s); i++ {
		parent := (i - 1) / 2
		newNode := NewNode(s[i], nil, nil)
		if s[i] == nullVal {
			continue
		}
		if i % 2 != 0 { // left node
			nodeSlice[parent].Left = newNode
		} else {
			nodeSlice[parent].Right = newNode
		}

		nodeSlice = append(nodeSlice, newNode)
	}
	return nodeSlice[0]
}

func NewNode[T comparable](val T, left, right *Node[T]) *Node[T] {
	return &Node[T]{
		Val: val,
		Left: left,
		Right: right,
	}
}
