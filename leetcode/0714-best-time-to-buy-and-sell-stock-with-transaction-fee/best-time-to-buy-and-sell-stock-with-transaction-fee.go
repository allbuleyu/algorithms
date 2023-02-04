package prob0714

func maxProfit(prices []int, fee int) int {
	return bottomUp(prices, fee)
}

func bottomUp(prices []int, fee int) int {
	var buy, sell int
	buy = -prices[0]

	for i := 1; i < len(prices); i++ {
		sell = max(sell, buy+prices[i]-fee)
		buy = max(buy, -prices[i]+sell)
	}

	return sell
}

var memory [][]int

func helper(prices []int, fee int) int {
	memory = make([][]int, 2)

	for i := 0; i < 2; i++ {
		memory[i] = make([]int, len(prices))
		for j := 0; j < len(prices); j++ {
			memory[i][j] = -1
		}
	}
	return upDown(prices, fee, 0, 0)
}

func upDown(prices []int, fee, x, y int) int {
	if x == len(prices) {
		return 0
	}

	if memory[y][x] != -1 {
		return memory[y][x]
	}

	res := 0
	if y == 0 {
		res = max(upDown(prices, fee, x+1, 0), -prices[x]+upDown(prices, fee, x+1, 1))
	} else {
		res = max(prices[x]-fee+upDown(prices, fee, x+1, 0), upDown(prices, fee, x+1, 1))
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
