package prob0091

func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	return bottomUp(s)
}

// todo
func bottomUp(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	if s[0] == '0' {
		dp[1] = 0
	}
	for i := 2; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] < '7') {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

var memory []int

func helperUpDown(s string) int {
	memory = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		memory[i] = -1
	}

	return upDown(s, 0)
}

func upDown(s string, k int) int {
	if k == len(s) {
		return 1
	}

	if memory[k] > -1 {
		return memory[k]
	}

	if s[k] == '0' {
		return 0
	}

	if k == len(s)-1 {
		return 0
	}

	ans := upDown(s, k+1)
	if s[k] == '1' || (s[k] == '2' && s[k+1] < '7') {
		ans += upDown(s, k+2)
	}

	memory[k] = ans
	return ans
}
