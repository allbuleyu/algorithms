package prob1155

import "math"

func numRollsToTarget(n int, k int, target int) int {
	if n*k < target {
		return 0
	}

	return bottomUpOptimization(n, k, target)
}

func bottomUpOptimization(n, k, target int) int {
	dp := make([]int, target+1)
	dpTmp := make([]int, target+1)
	dp[target] = 1

	for diceIndex := n - 1; diceIndex >= 0; diceIndex-- {
		for curSum := 0; curSum <= target; curSum++ {
			ways := 0
			for i := 1; i <= min(k, target-curSum); i++ {

				ways = (ways + dp[curSum+i]) % (int(math.Pow(10, 9)) + 7)
			}

			dpTmp[curSum] = ways
		}

		dp = dpTmp
	}

	return dp[0]
}

func bottomUp(n, k, target int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
	}

	dp[n][target] = 1
	for diceIndex := n - 1; diceIndex >= 0; diceIndex-- {
		for curSum := 0; curSum <= target; curSum++ {
			ways := 0
			for i := 1; i <= min(k, target-curSum); i++ {

				ways = (ways + dp[diceIndex+1][curSum+i]) % (int(math.Pow(10, 9)) + 7)
			}

			dp[diceIndex][curSum] = ways
		}
	}

	return dp[0][0]
}

var memory [][]int

func helperTopDown(n, k, target int) int {
	memory = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memory[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			memory[i][j] = -1
		}
	}

	return topDown(n, k, target)
}

func topDown(n, k, target int) int {
	if n == 0 && target == 0 {
		return 1
	}

	if n == 0 || target <= 0 {
		return 0
	}

	if memory[n][target] != -1 {
		return memory[n][target]
	}

	res := 0
	for i := 1; i <= k; i++ {
		res = (res + topDown(n-1, k, target-i)) % (int(math.Pow(10, 9)) + 7)
	}
	memory[n][target] = res

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
