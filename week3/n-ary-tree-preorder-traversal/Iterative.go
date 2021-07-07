package main

// 迭代法遍历
func (t *traversal) preorderIterative(root *Node) {
	if root == nil {
		return
	}
	stack := newStack()
	stack.push(root)
	for {
		node := stack.pop()
		if len(node.Children) > 0 {
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack.push(node.Children[i])
			}
		}
		t.list = append(t.list, node.Val)
		if stack.empty() {
			return
		}
	}
}

type stack struct {
	list []*Node
	cur  int
}

func newStack() *stack {
	return &stack{
		list: make([]*Node, 50),
	}
}

func (t *stack) push(node *Node) {
	if len(t.list) == t.cur+1 { // increase list size
		l := make([]*Node, len(t.list)*2)
		copy(l, t.list)
		t.list = l
	}
	t.list[t.cur] = node
	t.cur++
}

func (t *stack) pop() *Node {
	t.cur--
	v := t.list[t.cur]
	if len(t.list)/4 > t.cur+1 { // reduce list size. may not be necessary
		l := make([]*Node, len(t.list)/2)
		copy(l, t.list)
		t.list = l
	}
	return v
}

func (t *stack) empty() bool {
	return t.cur == 0
}
