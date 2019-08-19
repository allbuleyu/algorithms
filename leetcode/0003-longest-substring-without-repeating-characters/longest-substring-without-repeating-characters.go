package prob0003

func lengthOfLongestSubstring(s string) int {
	return slideWindow3(s)
}

func slideWindow(s string) int {
	var ans, i, j int
	hs := make(map[uint8]bool)

	for i < len(s) && j < len(s) {
		if _, ok := hs[s[j]]; !ok {
			hs[s[j]] = true
			j++
			ans = max(ans, j-i)

		}else {
			delete(hs, s[i])
			i++
		}
	}

	return ans
}

func slideWindow2(s string) int {
	var ans, i, j int
	hs := make(map[uint8]int)

	for i < len(s) && j < len(s) {
		if v, ok := hs[s[j]]; ok {
			i = max(v, i)
		}

		hs[s[j]] = j+1
		j++
		ans = max(ans, j-i)
	}

	return ans
}

// ASCII 256 字符最多
func slideWindow3(s string) int {
	var tb [256]int
	var ans, i, j int

	for ; j < len(s); j++ {
		i = max(i, tb[s[j]])
		ans = max(ans, j+1-i)

		tb[s[j]] = j+1
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}