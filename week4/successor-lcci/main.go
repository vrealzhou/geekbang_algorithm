/*
设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。

如果指定节点没有对应的“下一个”节点，则返回null。

https://leetcode-cn.com/problems/successor-lcci/
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	// has right sub-tree, leftest leaf of right sub-tree
	if p.Right != nil {
		node := p.Right
		for node != nil && node.Left != nil {
			node = node.Left
		}
		return node
	}
	// store parents
	node := root
	stack := make([]*TreeNode, 0)
	for node != nil && node.Val != p.Val {
		stack = append(stack, node)
		if node.Val > p.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	if len(stack) == 0 {
		return nil
	}
	top := len(stack) - 1
	n := p
	for top >= 0 {
		// be left node of parent, parent
		if stack[top].Left != nil && stack[top].Left.Val == n.Val {
			return stack[top]
		}
		// be right node of parent, parent's parent
		n = stack[top]
		top--
	}
	return nil
}
