package main

// 迭代法遍历
func (t *traversal) levelOrderIterative(root *Node) {
	if root == nil {
		return
	}
	t.list = append(t.list, []int{})
	q := newQueue()
	q.push(root)
	q.push(nil)
	level := 0
	for {
		node := q.pop()
		if node == nil {
			if q.empty() {
				return
			}
			q.push(nil)
			level++
			t.list = append(t.list, []int{})
		} else {
			t.list[level] = append(t.list[level], node.Val)
			for _, child := range node.Children {
				q.push(child)
			}
		}
	}
}

type queue struct {
	list []*Node
	tail int
}

func newQueue() *queue {
	return &queue{
		list: make([]*Node, 50),
	}
}

func (t *queue) push(node *Node) {
	if len(t.list) == t.tail+1 { // increase list size
		l := make([]*Node, len(t.list)*2)
		copy(l, t.list)
		t.list = l
	}
	t.list[t.tail] = node
	t.tail++
}

func (t *queue) pop() *Node {
	t.tail--
	v := t.list[0]
	t.list = t.list[1:]
	return v
}

func (t *queue) empty() bool {
	return t.tail == 0
}
