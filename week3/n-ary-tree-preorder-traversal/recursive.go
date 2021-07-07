/*
https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
给定一个 N 叉树，返回其节点值的 前序遍历 。

N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

提示：

    N 叉树的高度小于或等于 1000
    节点总数在范围 [0, 10^4] 内

*/
package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	t := &traversal{
		list: make([]int, 0),
	}
	t.preorderIterative(root)
	return t.list
}

type traversal struct {
	list []int
}

// 递归法遍历
func (t *traversal) preorderRecursive(node *Node) {
	if node == nil {
		return
	}
	t.list = append(t.list, node.Val)
	for _, child := range node.Children {
		t.preorderRecursive(child)
	}
}

type Node struct {
	Val      int
	Children []*Node
}
