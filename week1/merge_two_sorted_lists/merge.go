/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
package merge

import (
	"github.com/vrealzhou/geekbang_algorithm/internal/data"
)

type ListNode = data.IntListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	node := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			node.Next = l2
			l2 = l2.Next
		} else if l2 == nil {
			node.Next = l1
			l1 = l1.Next
		} else if l1.Val <= l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	return head.Next
}
