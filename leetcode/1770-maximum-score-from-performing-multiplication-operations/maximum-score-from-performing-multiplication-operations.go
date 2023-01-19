package prob1770

func maximumScore(nums []int, multipliers []int) int {
	return helper(nums, multipliers)
}

func helper(nums, muls []int) int {
	return upDown(nums, muls, 0, 0)
}

func upDown(nums, muls []int, l, i int) int {
	if i == len(muls) {
		return 0
	}

	n := len(nums)

	resLeft := upDown(nums, muls, l+1, i+1)
	resRight := upDown(nums, muls, l, i+1)

	xl, xr := nums[l]*muls[i], nums[n-1-(i-l)]*muls[i]

	return max(resLeft+xl, resRight+xr)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
