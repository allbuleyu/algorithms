package prob0152

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return dynamicProgramming(nums)
}

func dynamicProgramming(nums []int) int {
	dp_max := make([]int, len(nums))
	dp_min := make([]int, len(nums))

	dp_max[0] = nums[0]
	dp_min[0] = nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		dp_max[i] = max(nums[i], dp_max[i-1] * nums[i], dp_min[i-1] * nums[i])
		dp_min[i] = min(nums[i], dp_max[i-1] * nums[i], dp_min[i-1] * nums[i])

		if res < dp_max[i] {
			res = dp_max[i]
		}
	}

	return res
}

func wrongAnswer(nums []int) int {
	dp := make([]int, len(nums))

	dp[0] = nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			dp[i] = 0
		}else if dp[i-1] == 0 {
			dp[i] = nums[i]
		}else {
			// wrong test 3,-1,4 还是考虑不够全面,漏掉一种情况
			dp[i] = dp[i-1] * nums[i]
		}

		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
}

func max(a, b, c int) int {
	m := a
	if m < b {
		m = b
	}

	if m < c {
		m = c
	}

	return m
}

func min(a, b, c int) int {
	m := a
	if m > b {
		m = b
	}

	if m > c {
		m = c
	}

	return m
}