/*
 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parentsP := findParents(root, p)
	parentsQ := findParents(root, q)
	// printNodes(parentsP)
	// printNodes(parentsQ)
	i := len(parentsP) - 1
	j := len(parentsQ) - 1
	var parent *TreeNode
	for i >= 0 && j >= 0 && parentsP[i].Val == parentsQ[j].Val {
		parent = parentsP[i]
		i--
		j--
	}
	return parent
}

func printNodes(list []*TreeNode) {
	fmt.Print("[")
	for i, node := range list {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(node.Val)
	}
	fmt.Println("]")
}

func findParents(node, v *TreeNode) []*TreeNode {
	f := NewParentFinder(v)
	f.find(node)
	return f.parents
}

type ParentFinder struct {
	parents []*TreeNode
	v       *TreeNode
}

func NewParentFinder(v *TreeNode) *ParentFinder {
	return &ParentFinder{
		parents: make([]*TreeNode, 0),
		v:       v,
	}
}

func (p *ParentFinder) find(node *TreeNode) {
	if node == nil {
		return
	}
	if node.Val == p.v.Val {
		p.parents = append(p.parents, node)
		return
	}
	p.find(node.Left)
	if len(p.parents) > 0 {
		p.parents = append(p.parents, node)
		return
	}
	p.find(node.Right)
	if len(p.parents) > 0 {
		p.parents = append(p.parents, node)
		return
	}
}
