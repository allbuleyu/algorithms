package prob0014

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	return divideAndConquer(strs)
}

func lcs(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	lcs := strs[0]

	for i := 1; i < len(strs); i++ {
		n := len(lcs)

		if n > len(strs[i]) {
			n = len(strs[i])
		}

		lcs = lcs[:n]

		for j := 0; j < n; j++ {
			if lcs[j] != strs[i][j] {
				lcs = lcs[:j]
				break
			}


			if lcs == "" {
				return ""
			}
		}


	}

	return lcs
}

func divideAndConquer(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}



	return ""
}