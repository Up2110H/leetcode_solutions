package bitManipulation

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	count := 0
	countOfPseudo(root, [10]int{}, 0, &count)
	return count
}

func countOfPseudo(node *TreeNode, counter [10]int, countOfOdd int, count *int) {
	counter[node.Val]++

	countOfOdd += counter[node.Val]&1 - (counter[node.Val]+1)&1

	if node.Left == nil && node.Right == nil && countOfOdd <= 1 {
		*count++
	}

	if node.Left != nil {
		countOfPseudo(node.Left, counter, countOfOdd, count)
	}

	if node.Right != nil {
		countOfPseudo(node.Right, counter, countOfOdd, count)
	}

	counter[node.Val]--
	countOfOdd += counter[node.Val]&1 - (counter[node.Val]+1)&1
}
