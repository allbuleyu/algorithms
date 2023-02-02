package prob0309

func maxProfit(prices []int) int {
	return bottomUp(prices)
}

// Now, I'm get significant through for dynamic programming.
// Congratulations,Hyl! Keep on fighting.
func bottomUp(prices []int) int {
	n := len(prices)
	var none, buy, sell int
	buy = -prices[0]
	for i := 1; i < n; i++ {
		none1 := max(none, sell)
		buy1 := max(-prices[i]+none, buy)
		sell = buy + prices[i]

		none, buy = none1, buy1
	}

	return max(none, sell)
}

var memory [][]int

func helperUpDown(prices []int) int {
	memory = make([][]int, 3)
	for i := 0; i < 3; i++ {
		memory[i] = make([]int, len(prices))

		for j := 0; j < len(prices); j++ {
			memory[i][j] = -1
		}
	}

	return upDown(prices, 0, 0)
}

// y 0 未持有,1 持有股票, 2 冷却
func upDown(prices []int, x, y int) int {
	if x == len(prices) {
		return 0
	}

	if memory[y][x] != -1 {
		return memory[y][x]
	}

	res := 0
	if y == 0 {
		res = max(upDown(prices, x+1, 0), -prices[x]+upDown(prices, x+1, 1))
	} else if y == 1 {
		res = max(upDown(prices, x+1, 1), prices[x]+upDown(prices, x+1, 2))
	} else {
		res = upDown(prices, x+1, 0)
	}
	memory[y][x] = res

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
