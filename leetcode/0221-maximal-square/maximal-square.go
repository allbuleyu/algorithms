package prob0221

func maximalSquare(matrix [][]byte) int {
	edge := bottomUp(matrix)

	return edge * edge
}

func bottomUp(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([]int, n+1)
	edge := 0
	diagonal := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if matrix[i-1][j-1]-'0' == 1 {
				dp[j] = min(dp[j], min(dp[j-1], diagonal)) + 1

				if dp[j] > edge {
					edge = dp[j]
				}
			} else {
				dp[j] = 0
			}

			diagonal = tmp
		}
	}

	return edge
}

var memory [][]int

func dfsHelper(matrix [][]byte) int {
	memory = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		sub := make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			sub[j] = -1
		}
		memory[i] = sub
	}
	return dfs(matrix, len(matrix)-1, len(matrix[0])-1)
}

func dfs(matrix [][]byte, i, j int) (ans int) {
	if i == 0 || j == 0 {
		ans = int(matrix[i][j] - '0')
		return
	}

	if memory[i][j] != -1 {
		return memory[i][j]
	}

	if matrix[i][j] == 0 {
		memory[i][j] = 0
		return 0
	}

	ans = min(dfs(matrix, i-1, j-1), min(dfs(matrix, i-1, j), dfs(matrix, i, j-1)))
	memory[i][j] = ans

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
