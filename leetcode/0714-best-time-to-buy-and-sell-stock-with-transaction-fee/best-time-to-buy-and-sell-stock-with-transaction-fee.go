package prob0714

func maxProfit(prices []int, fee int) int {
	var none, buy, sell int
	buy = prices[0]
	// todo: 状态转变回去再写
	for i := 1; i < len(prices); i++ {
		none1 := max(none, max(sell, buy+prices[i]))
		buy1 := max(-prices[i]+none, max(buy, -prices[i]+sell))
		sell = buy + prices[i]

		none, buy = none1, buy1
	}

	return sell
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
