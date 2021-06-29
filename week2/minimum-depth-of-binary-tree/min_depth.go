/*
	https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

	给定一个二叉树，找出其最小深度。
	最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
	说明：叶子节点是指没有子节点的节点。

	提示：
    - 树中节点数的范围在 [0, 105] 内
    - -1000 <= Node.val <= 1000

*/

package minimumdepth

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	leftDepth := int(math.MaxInt32)
	rightDepth := int(math.MaxInt32)
	if root.Left != nil {
		leftDepth = 1 + minDepth(root.Left)
	}
	if root.Right != nil {
		rightDepth = 1 + minDepth(root.Right)
	}
	if leftDepth < rightDepth {
		return leftDepth
	}
	return rightDepth
}
