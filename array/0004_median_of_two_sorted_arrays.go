package array

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	b := findKNum(nums1, nums2, 0, 0, len(nums1), len(nums2), length/2+1)

	if length%2 == 0 {
		return float64(b+findKNum(nums1, nums2, 0, 0, len(nums1), len(nums2), length/2)) / 2
	}

	return float64(b)
}

func findKNum(n1 []int, n2 []int, o1 int, o2 int, l1 int, l2 int, k int) int {
	if l1 > l2 {
		return findKNum(n2, n1, o2, o1, l2, l1, k)
	}

	if l1 == 0 {
		return n2[o2+k-1]
	}

	if k == 1 {
		return int(math.Min(float64(n1[o1]), float64(n2[o2])))
	}

	p := int(math.Min(float64(l1), float64(k/2)))
	q := k - p

	if n1[o1+p-1] == n2[o2+q-1] {
		return n1[o1+p-1]
	}

	if n1[o1+p-1] < n2[o2+q-1] {
		return findKNum(n1, n2, o1+p, o2, l1-p, l2, k-p)
	}

	return findKNum(n1, n2, o1, o2+q, l1, l2-q, k-q)
}
