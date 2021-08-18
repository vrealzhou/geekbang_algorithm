package data

import (
	"fmt"
	"io"
	"strconv"
)

type AVLTree struct {
	root *AVLNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

type AVLKey interface {
	Compare(AVLKey) int
	String() string
}

type AVLNode struct {
	Data    AVLKey
	Count   int // 对于重复数据，增加count
	Balance int
	Height  int
	Left    *AVLNode
	Right   *AVLNode
}

func (n *AVLNode) Print(w io.Writer, ch rune) {
	fmt.Fprintf(w, "%c:%s, h:%d, bal:%d, count:%d\n", ch, n.Data.String(), n.Height, n.Balance, n.Count)
}

func (t *AVLTree) Max() AVLKey {
	node := t.root
	for node != nil {
		if node.Right == nil {
			return node.Data
		}
		node = node.Right
	}
	return nil
}

func (t *AVLTree) Find(d AVLKey) AVLKey {
	if t.root == nil {
		return nil
	}
	n := t.find(d, t.root)
	if n == nil {
		return nil
	}
	return n.Data
}

func (t *AVLTree) find(d AVLKey, node *AVLNode) *AVLNode {
	comp := d.Compare(node.Data)
	if comp == 0 {
		return node
	} else if comp < 0 {
		if node.Left == nil {
			return nil
		}
		return t.find(d, node.Left)
	}
	if node.Right == nil {
		return nil
	}
	return t.find(d, node.Right)
}

func (t *AVLTree) newNode(d AVLKey) *AVLNode {
	return &AVLNode{
		Data:    d,
		Count:   1,
		Balance: 0,
		Height:  1,
	}
}
func (t *AVLTree) Insert(d AVLKey) {
	if t.root == nil {
		t.root = &AVLNode{
			Data:    d,
			Count:   1,
			Balance: 0,
			Height:  1,
		}
		return
	}
	t.root = t.insert(d, t.root)
}

func (t *AVLTree) insert(d AVLKey, node *AVLNode) *AVLNode {
	comp := d.Compare(node.Data)
	if comp == 0 {
		node.Count++
		return node
	} else if comp < 0 {
		if node.Left == nil {
			node.Left = t.newNode(d)
			fixHeightAndBalance(node)
			return node
		}
		n := t.insert(d, node.Left)
		node.Left = n
		fixHeightAndBalance(node)
		return tryBalance(node)
	}
	if node.Right == nil {
		node.Right = t.newNode(d)
		fixHeightAndBalance(node)
		return node
	}
	n := t.insert(d, node.Right)
	node.Right = n
	fixHeightAndBalance(node)
	return tryBalance(node)
}

func (t *AVLTree) Successor(d AVLKey) AVLKey {
	p := t.find(d, t.root)
	if p == nil {
		return nil
	}
	p = t.getSuccessor(p)
	if p == nil {
		return nil
	}
	return p.Data
}

func (t *AVLTree) getSuccessor(p *AVLNode) *AVLNode {
	// has right sub-tree, leftest leaf of right sub-tree
	if p.Right != nil {
		node := p.Right
		for node != nil && node.Left != nil {
			node = node.Left
		}
		return node
	}
	// find parents
	node := t.root
	stack := make([]*AVLNode, 0)
	for node != nil {
		comp := node.Data.Compare(p.Data)
		if comp != 0 {
			stack = append(stack, node)
			if comp > 0 {
				node = node.Left
			} else {
				node = node.Right
			}
		} else {
			break
		}
	}
	if len(stack) == 0 {
		return nil
	}
	top := len(stack) - 1
	n := p
	for top >= 0 {
		// be left node of parent, parent
		if stack[top].Left != nil && stack[top].Left.Data.Compare(n.Data) == 0 {
			if top > 1 {
				return stack[top]
			}
			return stack[top]
		}
		// be right node of parent, parent's parent
		n = stack[top]
		top--
	}
	return nil
}

func (t *AVLTree) Predecessor(d AVLKey) AVLKey {
	p := t.find(d, t.root)
	if p == nil {
		return nil
	}
	p = t.getPredecessor(p)
	if p == nil {
		return nil
	}
	return p.Data
}

func (t *AVLTree) getPredecessor(p *AVLNode) *AVLNode {
	// has left sub-tree, rightest leaf of right sub-tree
	if p.Left != nil {
		node := p.Left
		for node != nil && node.Right != nil {
			node = node.Right
		}
		return node
	}
	// find parents
	node := t.root
	stack := make([]*AVLNode, 0)
	for node != nil {
		comp := node.Data.Compare(p.Data)
		if comp != 0 {
			stack = append(stack, node)
			if comp > 0 {
				node = node.Left
			} else {
				node = node.Right
			}
		} else {
			break
		}
	}
	if len(stack) == 0 {
		return nil
	}
	top := len(stack) - 1
	n := p
	for top >= 0 {
		// be right node of parent, parent
		if stack[top].Right != nil && stack[top].Right.Data.Compare(n.Data) == 0 {
			if top > 1 {
				return stack[top]
			}
			return stack[top]
		}
		// be right node of parent, parent's parent
		n = stack[top]
		top--
	}
	return nil
}

func (t *AVLTree) Remove(d AVLKey) bool {
	if t.root == nil {
		return false
	}
	n, removed := t.remove(d, t.root, false)
	t.root = n
	return removed
}

func (t *AVLTree) remove(d AVLKey, node *AVLNode, ignoreCount bool) (*AVLNode, bool) {
	comp := d.Compare(node.Data)
	if comp == 0 {
		if !ignoreCount {
			node.Count--
			if node.Count > 0 {
				return node, true
			}
		}
		if node.Left != nil && node.Right != nil {
			successor := t.getSuccessor(node)
			n, _ := t.remove(successor.Data, node.Right, true)
			successor.Right = n
			successor.Left = node.Left
			fixHeightAndBalance(successor)
			return tryBalance(successor), true
		} else if node.Left != nil {
			return node.Left, true
		} else if node.Right != nil {
			return node.Right, true
		} else {
			return nil, true
		}
	} else if comp < 0 {
		if node.Left == nil {
			return node, false
		}
		n, removed := t.remove(d, node.Left, ignoreCount)
		if removed {
			node.Left = n
			fixHeightAndBalance(node)
			return tryBalance(node), removed
		}
		return node, false
	} else {
		if node.Right == nil {
			return node, false
		}
		n, removed := t.remove(d, node.Right, ignoreCount)
		if removed {
			node.Right = n
			fixHeightAndBalance(node)
			// printTree(os.Stdout, t.root, 0, 'M')
			return tryBalance(node), removed
		}
		return node, false
	}
}

func fixHeightAndBalance(n *AVLNode) {
	max := 0
	bal := 0
	if n.Left != nil {
		max = n.Left.Height
		bal = max - 0
	}
	if n.Right != nil {
		bal = max - n.Right.Height
		if max < n.Right.Height {
			max = n.Right.Height
		}
	}
	n.Height = max + 1
	n.Balance = bal
}

func tryBalance(n *AVLNode) *AVLNode {
	if n.Balance >= -1 && n.Balance <= 1 {
		return n
	}
	if n.Balance > 1 {
		if n.Left.Balance >= 0 { // LL
			b := n.Left
			n.Left = b.Right
			b.Right = n
			fixHeightAndBalance(n)
			fixHeightAndBalance(b)
			return b
		} else { // LR
			a := n.Left
			b := a.Right
			a.Right = b.Left
			b.Left = a
			fixHeightAndBalance(a)
			n.Left = b.Right
			b.Right = n
			fixHeightAndBalance(n)
			fixHeightAndBalance(b)
			return b
		}
	} else {
		if n.Right.Balance <= 0 { // RR
			b := n.Right
			n.Right = b.Left
			b.Left = n
			fixHeightAndBalance(n)
			fixHeightAndBalance(b)
			return b
		} else { // RL
			a := n.Right
			b := a.Left
			a.Left = b.Right
			b.Right = a
			fixHeightAndBalance(a)
			n.Right = b.Left
			b.Left = n
			fixHeightAndBalance(n)
			fixHeightAndBalance(b)
			return b
		}
	}
}

type IntKey int

func (i IntKey) Compare(k AVLKey) int {
	i1, ok := k.(IntKey)
	if !ok {
		panic(fmt.Sprintf("Incorrect value %v\n", k))
	}
	return int(i) - int(i1)
}

func (i IntKey) String() string {
	return strconv.Itoa(int(i))
}

func printTree(w io.Writer, node *AVLNode, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	node.Print(w, ch)
	printTree(w, node.Left, ns+2, 'L')
	printTree(w, node.Right, ns+2, 'R')
}
