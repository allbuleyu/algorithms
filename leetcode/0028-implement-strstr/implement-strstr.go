package prob0028

func strStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)
	if m < n {
		return -1
	}

	if needle == "" {
		return 0
	}

	start := -1

	// 时间复杂负 (m - n ) * n
	HS:
	for i := 0; i < m; i++ {
		if i + n > m {
			break HS
		}

		tmp := i
		for j := 0; j < n;  {
			if haystack[tmp] != needle[j] {
				continue HS
			}

			tmp++
			j++
		}

		return i
	}

	return start
}