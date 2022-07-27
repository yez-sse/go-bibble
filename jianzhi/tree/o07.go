package tree

import "go-learning/struct"

func buildTree(preorder, inorder []int) *_struct.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &_struct.TreeNode{Val: preorder[0]}
	i := 0 // 这里可以用一个map，初始化之后直接找，当然两者都需要树中没有重复数字
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	} // 每次inorder都会变！preorder也是！
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root
}
