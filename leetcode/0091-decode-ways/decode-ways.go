package prob0091

func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	return bottomUp(s)
}

func bottomUp(s string) int {
	n := len(s)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		if s[i-1] == '0' && s[i] == '0' {
			return 0
		}

		if s[i-1] == '1' || (s[i-1] == '2' && s[i] < '7') {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[n-1]
}

func helperUpDown(s string) int {
	return upDown(s, len(s)-1)
}

func upDown(s string, k int) int {
	if k == 0 {
		return 1
	}

	if s[k-1] == '0' && s[k] == '0' {
		return 0
	}

	if s[k-1] == '1' || (s[k-1] == '2' && s[k] < '7') {
		return upDown(s, k-1) + 1
	}

	return upDown(s, k-1)
}
