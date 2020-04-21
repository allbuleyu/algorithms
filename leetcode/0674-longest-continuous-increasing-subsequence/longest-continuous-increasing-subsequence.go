package prob0674

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return slidingWindow(nums)
}

func slidingWindow(nums []int) int {
	var res, cur int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] >= nums[i] {
			cur = i
		}
		res = max(res, i-cur+1)
	}

	return res
}

func normal(nums []int) int {
	maxLen := 1
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cur++
		}else {
			cur = 1
		}

		maxLen = max(maxLen, cur)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minimumSubSquence(nums []int, target int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums) && target > 0; i++ {
		for target - nums[i] < 0 && i < len(nums) -1 {
			if target - nums[i+1] > 0 {
				break
			}

			i++
		}

		target -= nums[i]
		res = append(res, nums[i])
	}

	return res
}