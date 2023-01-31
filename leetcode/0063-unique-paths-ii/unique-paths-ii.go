package prob0063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}

		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[0] = 0
		}

		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[n-1]
}
