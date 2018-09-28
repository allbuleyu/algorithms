package prob0121


func maxProfit(prices []int) int {
	n := len(prices)

	var tmp, res int
	for i := 1; i < n; i++ {
		tmp += prices[i] - prices[i-1]
		if tmp < 0 {
			tmp = 0		// 不用考虑为负的情况
		}

		if tmp > res {
			res = tmp
		}
	}

	return res
}
