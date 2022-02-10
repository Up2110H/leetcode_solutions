package stack

func insert(node *TreeNode, val int) {
	if val < node.Val {
		if node.Left != nil {
			insert(node.Left, val)
		} else {
			node.Left = &TreeNode{Val: val}
		}
	} else {
		if node.Right != nil {
			insert(node.Right, val)
		} else {
			node.Right = &TreeNode{Val: val}
		}
	}
}

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{Val: preorder[0]}

	for i := 1; i < len(preorder); i++ {
		insert(root, preorder[i])
	}

	return root
}
