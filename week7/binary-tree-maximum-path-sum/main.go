/*
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := int(math.MinInt32)
	sumRoot, _ := dfs(root, &maxSum)
	return max(maxSum, sumRoot)
}

func dfs(n *TreeNode, maxSum *int) (int, int) {
	// fmt.Printf("maxSum: %d", maxSum)
	sumn := 0 // 返回sum的时候需要考虑左右子树max值为负数的情况，如果是负数就不带
	maxn := 0 // 返回max的时候需要考虑左右子树max值为负数的情况，如果是负数就不带
	if n.Left != nil {
		suml, maxl := dfs(n.Left, maxSum)
		// fmt.Printf(", suml: %d, maxl: %d", suml, maxl)
		*maxSum = max(*maxSum, suml)
		if maxl > 0 {
			sumn = maxl
			maxn = maxl
		}
	}
	if n.Right != nil {
		sumr, maxr := dfs(n.Right, maxSum)
		// fmt.Printf(", sumr: %d, maxr: %d", sumr, maxr)
		*maxSum = max(*maxSum, sumr)
		if maxr > 0 {
			sumn = sumn + maxr
			maxn = max(maxn, maxr)
		}
	}
	// fmt.Printf(", n.Val: %d, maxSum2: %d\n", n.Val, maxSum)
	return sumn + n.Val, maxn + n.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
