package prob0188

func maxProfit(k int, prices []int) int {
	return bottomUp(k, prices)
}

// todo
//func Merge(k int, prices []int) int {
//
//}

func bottomUp(k int, prices []int) int {
	n := len(prices)
	dp := make([][][2]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][2]int, k+1)
	}

	for i := n - 1; i >= 0; i-- {
		for x := k - 1; x >= 0; x-- {
			for h := 0; h < 2; h++ {
				doNothing := dp[i+1][x][h]
				doSomething := 0
				if h == 1 {
					doSomething = prices[i] + dp[i+1][x+1][0]
				} else {
					doSomething = -prices[i] + dp[i+1][x][1]
				}

				dp[i][x][h] = max(doSomething, doNothing)
			}
		}
	}

	return dp[0][0][0]
}

func helperUpDown(k int, prices []int) int {
	return upDown(prices, 0, k, 0)
}

func upDown(prices []int, i, k, isHold int) int {
	if k == 0 || i == len(prices) {
		return 0
	}

	doNoting := upDown(prices, i+1, k, isHold)

	var doSomething int
	if isHold == 1 {
		doSomething = prices[i] + upDown(prices, i+1, k-1, 0)
	} else {
		doSomething = -prices[i] + upDown(prices, i+1, k, 1)
	}

	return max(doSomething, doNoting)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
