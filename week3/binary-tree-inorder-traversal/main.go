package main

/**
给定一个二叉树的根节点 root ，返回它的 中序 遍历。

示例 1：

输入：root = [1,null,2,3]
输出：[1,3,2]

示例 2：

输入：root = []
输出：[]

示例 3：

输入：root = [1]
输出：[1]

示例 4：

输入：root = [1,2]
输出：[2,1]

示例 5：

输入：root = [1,null,2]
输出：[1,2]



提示：

    树中节点数目在范围 [0, 100] 内
    -100 <= Node.val <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func inorderTraversal(root *TreeNode) []int {
	t := &traversal{
		list: make([]int, 0),
	}
	t.inorderTraversal(root)
	return t.list
}

type traversal struct {
	list []int
}

func (t *traversal) inorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		t.inorderTraversal(node.Left)
	}
	t.list = append(t.list, node.Val)
	if node.Right != nil {
		t.inorderTraversal(node.Right)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
