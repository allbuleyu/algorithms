package prob0005

func longestPalindrome(s string) string {

	return dynamicProgramming(s)
}


func expandAroundCenter(s string) string {
	var res string

	getPalindromic := func(start, end int) {
		for ; start >=0 && end < len(s); {
			if s[start] != s[end] {
				break
			}

			if len(s[start:end+1]) > len(res) {
				res = s[start:end+1]
			}

			start--
			end++
		}

	}

	for i := 0; i < len(s); i++ {
		getPalindromic(i,i)
		getPalindromic(i,i+1)
	}

	return res
}

func dynamicProgramming(s string) string {
	dp := make([][]bool, len(s))

	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ {
		for start, end := i, i; start >= 0 && end < len(s); {
			if s[start] != s[end] {
				break
			}

			dp[start][end] = true

			start--
			end++
		}


	}

	return ""
}