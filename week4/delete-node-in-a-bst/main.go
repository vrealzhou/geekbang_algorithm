/*
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

    首先找到需要删除的节点；
    如果找到了，删除它。

说明： 要求算法时间复杂度为 O(h)，h 为树的高度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/delete-node-in-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Right != nil {
			root.Val = successor(root).Val
			root.Right = deleteNode(root.Right, root.Val)
		} else if root.Left != nil {
			root.Val = predecessor(root).Val
			root.Left = deleteNode(root.Left, root.Val)
		} else {
			root = nil
		}
	}
	return root
}

func successor(root *TreeNode) *TreeNode {
	n := root.Right
	for n.Left != nil {
		n = n.Left
	}
	return n
}

func predecessor(root *TreeNode) *TreeNode {
	n := root.Left
	for n.Right != nil {
		n = n.Right
	}
	return n
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
