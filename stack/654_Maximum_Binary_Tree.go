package stack

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMax(nums []int, l, r int) *TreeNode {
	if l == r {
		return nil
	}

	maxVal := -1
	maxI := 0

	for i := l; i < r; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			maxI = i
		}
	}

	return &TreeNode{maxVal, getMax(nums, l, maxI), getMax(nums, maxI+1, r)}
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return getMax(nums, 0, len(nums))
}
