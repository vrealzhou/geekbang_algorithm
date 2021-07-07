/*
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	sb := strings.Builder{}
	q := newQueue()
	q.push(root)
	for {
		if q.empty() {
			break
		}
		sb.WriteRune(',')
		node := q.pop()
		if node == nil {
			sb.WriteString("null")
			continue
		}
		sb.WriteString(strconv.Itoa(node.Val))
		q.push(node.Left)
		q.push(node.Right)
	}
	sb.WriteRune(']')
	result := "[" + sb.String()[1:]
	return result
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	q := newQueue()
	data = data[1 : len(data)-1]
	nodes := strings.Split(data, ",")
	root := nodeFromValue(nodes[0])
	if root == nil {
		return root
	}
	q.push(root)
	for i := 1; i < len(nodes); i += 2 {
		left := nodes[i]
		right := nodes[i+1]
		nodeLeft := nodeFromValue(left)
		if nodeLeft != nil {
			q.push(nodeLeft)
		}
		nodeRight := nodeFromValue(right)
		if nodeRight != nil {
			q.push(nodeRight)
		}
		parent := q.pop()
		parent.Left = nodeLeft
		parent.Right = nodeRight
	}
	return root
}

func nodeFromValue(val string) *TreeNode {
	if val == "null" {
		return nil
	} else {
		intVal, _ := strconv.Atoi(val)
		return &TreeNode{
			Val: intVal,
		}
	}
}

type queue struct {
	list []*TreeNode
	tail int
}

func newQueue() *queue {
	return &queue{
		list: make([]*TreeNode, 50),
	}
}

func (t *queue) push(node *TreeNode) {
	if len(t.list) == t.tail+1 { // increase list size
		l := make([]*TreeNode, len(t.list)*2)
		copy(l, t.list)
		t.list = l
	}
	t.list[t.tail] = node
	t.tail++
}

func (t *queue) pop() *TreeNode {
	t.tail--
	if len(t.list) == 0 {
		return nil
	}
	v := t.list[0]
	t.list = t.list[1:]
	return v
}

func (t *queue) empty() bool {
	return t.tail == 0
}
