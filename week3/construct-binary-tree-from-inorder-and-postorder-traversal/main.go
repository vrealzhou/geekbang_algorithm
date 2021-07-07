package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	index := indexOf(inorder, postorder[len(postorder)-1])
	leftInOrder := inorder[:index]
	rightInOrder := inorder[index+1:]
	left := buildTree(leftInOrder, postorder[:index])
	right := buildTree(rightInOrder, postorder[index:len(postorder)-1])
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
