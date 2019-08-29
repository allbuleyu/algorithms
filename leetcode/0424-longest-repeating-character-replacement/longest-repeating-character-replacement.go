package prob0424

// confused
func characterReplacement(s string, k int) int {
	n := len(s)

	// maxCount 这个变量不知道是干嘛的...所以这里还没有搞懂
	var start, end, maxCount, maxLen int
	count := [26]int{}
	for ; end < n; end++ {
		count[s[end]-'A']++
		maxCount = max(maxCount, count[s[end]-'A'])
		for end-start+1-maxCount > k {
			count[s[start]-'A']--
			start++
		}

		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}