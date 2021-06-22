package data

type IntListNode struct {
	Val  int
	Next *IntListNode
}

func (n *IntListNode) Slice() []int {
	s := make([]int, 0)
	for node := n; node != nil; node = node.Next {
		s = append(s, node.Val)
	}
	return s
}

func IntList(items ...int) *IntListNode {
	head := &IntListNode{}
	node := head
	for _, item := range items {
		newNode := &IntListNode{
			Val: item,
		}
		node.Next = newNode
		node = newNode
	}
	return head.Next
}
