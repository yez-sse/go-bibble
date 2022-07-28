package lc1

import (
	_struct "go-learning/struct"
)

func removeElements(head *_struct.ListNode, val int) *_struct.ListNode {
	dummy := &_struct.ListNode{
		Val: -1,
	}
	dummy.Next = head
	p := dummy
	for p != nil {
		for p.Next != nil && p.Next.Val == val {
			p.Next = p.Next.Next
		}
		p = p.Next
	}
	return dummy.Next
}
//P206翻转链表
func reverseList(head *_struct.ListNode) *_struct.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &_struct.ListNode{Val: -1}
	p := head
	q := head.Next
	for p != nil {
		p.Next = dummy.Next
		dummy.Next = p
		p = q
		if q != nil {
			q = q.Next
		}
	}
	return dummy.Next
}
