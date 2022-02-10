package stack

type BSTIterator struct {
	root      *TreeNode
	node      *TreeNode
	ancestors []*TreeNode
}

func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	return BSTIterator{root, nil, make([]*TreeNode, 1)}
}

func (this *BSTIterator) Next() int {
	if this.node == nil {
		parent := this.root
		child := parent.Left

		for child != nil {
			this.ancestors = append(this.ancestors, parent)
			parent = child
			child = child.Left
		}

		this.node = parent
		return this.node.Val
	}

	if this.node.Right != nil {
		this.ancestors = append(this.ancestors, this.node)
		parent := this.node.Right
		child := parent.Left

		for child != nil {
			this.ancestors = append(this.ancestors, parent)
			parent = child
			child = child.Left
		}

		this.node = parent
		return this.node.Val
	}

	parent := this.ancestors[len(this.ancestors)-1]
	child := this.node

	for {
		this.ancestors = this.ancestors[:len(this.ancestors)-1]

		if parent.Left == child {
			this.node = parent
			return this.node.Val
		}

		child = parent
		parent = this.ancestors[len(this.ancestors)-1]
	}
}

func (this *BSTIterator) HasNext() bool {
	if this.node == nil || this.node.Right != nil {
		return true
	}

	i := 1
	parent := this.ancestors[len(this.ancestors)-i]
	child := this.node

	for parent != nil {
		if parent.Left == child {
			return true
		}

		i++
		child = parent
		parent = this.ancestors[len(this.ancestors)-i]
	}

	return false
}
