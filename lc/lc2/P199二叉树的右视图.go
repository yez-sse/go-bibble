package lc2

import _struct "go-learning/struct"

var res []int

func rightSideView(root *_struct.TreeNode) []int {
	res = []int{} //为什么这里一定要初始化一下呢？？
	dfs(root, 0)
	return res
}

func dfs(root *_struct.TreeNode, depth int) {
	if root == nil {
		return
	}
	if len(res) == depth {
		res = append(res, root.Val)
	}
	depth++
	dfs(root.Right, depth)
	dfs(root.Left, depth)
}

//bfs在go语言里应该怎么写

func rightSideView2(root *_struct.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	//queue = []*_struct.TreeNode{root}
	var queue []*_struct.TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		res = append(res, queue[0].Val)
		for length > 0 {
			length--
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			queue = queue[1:]
		}
	}
	return res
}
