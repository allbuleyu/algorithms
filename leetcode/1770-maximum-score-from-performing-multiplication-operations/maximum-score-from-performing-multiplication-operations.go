package prob1770

func maximumScore(nums []int, multipliers []int) int {
	return helper(nums, multipliers)
}

func helper(nums, muls []int) int {
	return f1(nums, muls, 0, len(nums)-1, 0)
}

func f1(nums, muls []int, l, r, i int) int {
	if i == len(muls) {
		return 0
	}

	resLeft := f1(nums, muls, l+1, r, i+1)
	resRight := f1(nums, muls, l, r-1, i+1)

	xl, xr := nums[l]*muls[i], nums[r]*muls[i]

	return max(resLeft+xl, resRight+xr)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
