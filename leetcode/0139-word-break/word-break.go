package prob0139

func wordBreak(s string, wordDict []string) bool {
	return bottomUp(s, wordDict)
}

func bottomUp(s string, wd []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[n] = true

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < len(wd); j++ {
			offset := len(wd[j])
			if i+offset > n || s[i:i+offset] != wd[j] {
				continue
			}

			dp[i] = dp[i+offset]
			if dp[i] {
				break
			}
		}
	}

	return dp[0]
}

var memory []int

func helpUpDown(s string, wordDict []string) bool {
	memory = make([]int, len(s))

	return upDown(s, 0, wordDict)
}

func upDown(s string, start int, wd []string) bool {
	if start == len(s) {
		return true
	}

	if memory[start] != 0 {
		return memory[start] == 1
	}

	ans := false
	for i := 0; i < len(wd); i++ {
		offset := len(wd[i])
		if start+offset > len(s) || wd[i] != s[start:start+offset] {
			continue
		}

		if upDown(s, start+offset, wd) {
			ans = true
			break
		}
	}

	memory[start] = -1
	if ans {
		memory[start] = 1
	}

	return ans
}
