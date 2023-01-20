package prob1770

func maximumScore(nums []int, multipliers []int) int {
	return bottomUp(nums, multipliers)
}

func bottomUp(nums, muls []int) int {
	m, n := len(nums), len(muls)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	for op := n - 1; op >= 0; op-- {
		for l := op; l >= 0; l-- {
			r := m - 1 - (op - l)
			dp[op][l] = max(muls[op]*nums[l]+dp[op+1][l+1], muls[op]*nums[r]+dp[op+1][l])
		}
	}

	return dp[0][0]
}

var memory [][]int

func helper(nums, muls []int) int {
	n := len(muls)
	memory = make([][]int, n)
	for i := 0; i < n; i++ {
		memory[i] = make([]int, n)
	}

	return upDown(nums, muls, 0, 0)
}

func upDown(nums, muls []int, l, i int) int {
	if i == len(muls) {
		return 0
	}

	if memory[i][l] > 0 {
		return memory[i][l]
	}

	n := len(nums)
	resLeft := upDown(nums, muls, l+1, i+1)
	resRight := upDown(nums, muls, l, i+1)

	xl, xr := nums[l]*muls[i], nums[n-1-(i-l)]*muls[i]
	memory[i][l] = max(resLeft+xl, resRight+xr)

	return memory[i][l]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
