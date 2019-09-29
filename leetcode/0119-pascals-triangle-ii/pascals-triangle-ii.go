package prob0119


func getRow(rowIndex int) []int {
	dp := make([]int, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		res := make([]int, i+1)
		dp[i] = 1

		for j := 1; j < i; j++ {
			res[j] = dp[j] + dp[j-1]
		}

		for j := 1; j < i; j++ {
			dp[j] = res[j]
		}
	}

	return dp
}