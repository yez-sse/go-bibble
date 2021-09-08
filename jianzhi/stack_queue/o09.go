package stack_queue

import "container/list"

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{ // 这里体会一下，struct中变量的定义，和变量初始化是分开的
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

func (this CQueue) DeleteHead() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	if this.stack2.Len() > 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int) // e.Value得到的是一个interface{}
	}
	return -1
}
