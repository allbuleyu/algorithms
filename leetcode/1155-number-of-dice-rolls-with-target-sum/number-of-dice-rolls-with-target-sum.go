package prob1155

import "math"

func numRollsToTarget(n int, k int, target int) int {
	if n*k < target {
		return 0
	}

	return helperTopDown(n, k, target)
}

// todo
func bottomUp(n, k, target int) int {

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
