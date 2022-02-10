package queue

func longestSubarray(nums []int, limit int) int {
	mx := make([]int, 0, len(nums))
	mn := make([]int, 0, len(nums))
	l := 0
	r := 0
	res := 1

	for r < len(nums) {
		for len(mx) > 0 && nums[mx[len(mx)-1]] <= nums[r] {
			mx = mx[:len(mx)-1]
		}
		for len(mn) > 0 && nums[mn[len(mn)-1]] >= nums[r] {
			mn = mn[:len(mn)-1]
		}

		mx = append(mx, r)
		mn = append(mn, r)

		for nums[mx[0]]-nums[mn[0]] > limit {
			p := mx[0]
			if p <= mn[0] {
				l = p + 1
				mx = mx[1:]
			}

			if p >= mn[0] {
				l = mn[0] + 1
				mn = mn[1:]
			}
		}

		r++
		if r-l > res {
			res = r - l
		}
	}

	return res
}
