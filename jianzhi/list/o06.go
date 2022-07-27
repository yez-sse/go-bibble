package list

import "container/list"
import "go-learning/struct"

// 翻转链表的方法
func reversePrint3(head *_struct.ListNode) []int {
	if head == nil {
		return nil
	}

	var newHead *_struct.ListNode
	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp // 后移一个，遍历
	}

	var ret []int // ret := []int{} 带上了初始化
	for newHead != nil {
		ret = append(ret, newHead.Val)
		newHead = newHead.Next
	}
	return ret
}

// 用递归的方法
func reversePrint2(head *_struct.ListNode) []int {
	if head == nil {
		return nil
	}
	return appendData(head)
}
func appendData(head *_struct.ListNode) []int {
	if head.Next != nil {
		// 这里的递归怕是比较难学会举一反三了
		list := appendData(head.Next)
		list = append(list, head.Val)
		return list
	}
	// head.Next == nil
	return []int{head.Val}
}

// 用栈的方法
func reversePrint(head *_struct.ListNode) []int {
	if head == nil {
		return nil
	}

	stack := list.New() // 返回的是*List
	for head != nil {   // 这里的for就是while
		stack.PushFront(head.Val)
		head = head.Next
	}
	var ret []int // 空的切片可以这样定义
	for e := stack.Front(); e != nil; e = e.Next() {
		ret = append(ret, e.Value.(int)) // 这句e.Value要注意学一下
	}
	return ret
}
