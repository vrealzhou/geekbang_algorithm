/*
https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

提示：

    树的高度不会超过 1000
    树的节点总数在 [0, 10^4] 之间

*/
package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func levelOrder(root *Node) [][]int {
	t := &traversal{
		list: make([][]int, 0),
	}
	t.levelOrderRecursive(0, root)
	return t.list
}

type traversal struct {
	list [][]int
}

// 递归法遍历
func (t *traversal) levelOrderRecursive(level int, node *Node) {
	if node == nil {
		return
	}
	if len(t.list) > level {
		t.list[level] = append(t.list[level], node.Val)
	} else {
		t.list = append(t.list, []int{node.Val})
	}
	for _, child := range node.Children {
		t.levelOrderRecursive(level+1, child)
	}
}

type Node struct {
	Val      int
	Children []*Node
}
