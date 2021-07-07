package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	index := indexOf(inorder, preorder[0])
	leftInOrder := inorder[:index]
	rightInOrder := inorder[index+1:]
	left := buildTree(preorder[1:len(leftInOrder)+1], leftInOrder)
	right := buildTree(preorder[len(leftInOrder)+1:], rightInOrder)
	root.Left = left
	root.Right = right
	return root
}

func indexOf(ary []int, val int) int {
	for i := 0; i < len(ary); i++ {
		if ary[i] == val {
			return i
		}
	}
	return -1
}
