package list

type ListNode struct {
	Val int
	// 要想有递归的定义，这里必须用指针，这也是指针的功能之一
	Next *ListNode
}
