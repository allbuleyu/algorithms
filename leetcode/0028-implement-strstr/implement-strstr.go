package prob0028


// 时间复杂度, m-n
func strStr(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	// 当hlen等于nlen的时候，需要i == 0
	for i := 0; i <= hlen-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}

	return -1
}

func strstr1(haystack string, needle string) int {
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

		if haystack[i:i+n] == needle {
			return i
		}

	}

	return start
}