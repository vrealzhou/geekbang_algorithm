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

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 分治做法
 */
func mergeKListsInDivide(lists []*ListNode) *ListNode {
	var head *ListNode
	var node *ListNode
	for {
		top := divideMin(lists, 0, len(lists))
		if top == -1 {
			break
		}
		if node != nil {
			node.Next = lists[top]
		}
		node = lists[top]
		if lists[top] != nil {
			lists[top] = lists[top].Next
		}
		if head == nil {
			head = node
		}
	}
	return head
}

func divideMin(lists []*ListNode, start, end int) int {
	if start == end {
		return -1
	} else if end-start == 1 {
		if lists[start] == nil {
			return -1
		}
		return start
	}
	split := (end - start) / 2
	if (end-start)%2 != 0 {
		split++
	}
	left := divideMin(lists, start, start+split)
	right := divideMin(lists, start+split, end)
	if left == -1 {
		return right
	} else if right == -1 {
		return left
	} else if lists[left].Val <= lists[right].Val {
		return left
	} else {
		return right
	}
}
