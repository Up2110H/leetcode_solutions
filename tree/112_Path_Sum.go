package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return dfs(root, 0, targetSum)
}

func dfs(node *TreeNode, sum, targetSum int) bool {
	sum += node.Val

	if node.Left == nil && node.Right == nil && sum == targetSum {
		return true
	}

	if node.Left != nil && dfs(node.Left, sum, targetSum) {
		return true
	}

	if node.Right != nil && dfs(node.Right, sum, targetSum) {
		return true
	}

	sum -= node.Val
	return false
}

/* 00:04 12.02.2022 - 00:16 12.02.2022 */
