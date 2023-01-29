package prob0518

func change(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	//for i := 1; i <= amount; i++ {
	//	for j := 0; j < len(coins); j++ {
	//		x := i - coins[j]
	//		if x < 0 {
	//			continue
	//		}
	//
	//		dp[i] += dp[x]
	//	}
	//}

	// 这个就是对的,上面的解法就是错的
	// example: coins:  []int{1, 2},amount: 3,
	// 上面的循环会统计[1,1,1],[1,2], [2,1],题目要求的是组合,所以上面的结果错误
	for j := 0; j < len(coins); j++ {
		for i := 1; i <= amount; i++ {
			x := i - coins[j]
			if x < 0 {
				continue
			}

			dp[i] += dp[x]
		}
	}

	// 标准dp解法: dp[i][j], i 代表第当前硬币,j 代表当前amount
	// dp[i][j] = dp[i-1][j] + dp[i][j-coins[i]] 状态转换方程. recurrence relation

	return dp[amount]
}
