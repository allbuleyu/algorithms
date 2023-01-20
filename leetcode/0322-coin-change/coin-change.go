package prob0322

func coinChange(coins []int, amount int) int {
	return bottomUp(coins, amount)
}

func bottomUp(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = maxVal
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] == maxVal {
		return -1
	}

	return dp[amount]
}

var memory []int
var maxVal = 1 << 31

func helpTopDown(coins []int, amount int) int {
	memory = make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		memory[i] = -2
	}

	topDown(coins, amount)
	if memory[amount] == maxVal {
		return -1
	}

	return memory[amount]
}

func topDown(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}

	if amount == 0 {
		return 0
	}

	if memory[amount] != -2 {
		return memory[amount]
	}

	minExchange := maxVal
	for i := 0; i < len(coins); i++ {
		x := topDown(coins, amount-coins[i])
		if x >= 0 && x < minExchange {
			minExchange = min(minExchange, x) + 1
		}
	}

	if minExchange == maxVal {
		memory[amount] = -1
	} else {
		memory[amount] = minExchange
	}

	return memory[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
