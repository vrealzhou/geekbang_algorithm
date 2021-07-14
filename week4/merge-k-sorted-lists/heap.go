/*

	给你一个链表数组，每个链表都已经按升序排列。

	请你将所有链表合并到一个升序链表中，返回合并后的链表。

	提示：

    k == lists.length
    0 <= k <= 10^4
    0 <= lists[i].length <= 500
    -10^4 <= lists[i][j] <= 10^4
    lists[i] 按 升序 排列
    lists[i].length 的总和不超过 10^4

	https://leetcode-cn.com/problems/merge-k-sorted-lists/

*/
package main

import "github.com/vrealzhou/geekbang_algorithm/internal/data"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 二叉堆做法
 */
func mergeKListsInDivide(lists []*ListNode) *ListNode {
	tops := data.NewHeap(len(lists), func(source, target interface{}) int {
		return target.(*ListNode).Val - source.(*ListNode).Val
	})
	var head *ListNode
	var node *ListNode
	for _, top := range lists { // init head
		tops.Add(top)
	}
	for {
		top := tops.Pop().(*ListNode)
		if top == nil {
			break
		}
		if head == nil {
			node = top
			head = node
		} else {
			node.Next = top
			node = top
		}
		if top.Next != nil {
			tops.Add(top.Next)
		}
	}
	return head
}
