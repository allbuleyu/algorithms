package prob0028

// 时间复杂度
func strStr(haystack string, needle string) int {
	m,n := len(haystack), len(needle)
	if n == 0 {
		return 0
	}

	if m < n {
		return -1
	}

	return indexRabinKarp(haystack, needle)
}

// primeRK is the prime base used in Rabin-Karp algorithm.
const primeRK = 16777619

// hashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func hashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

func indexRabinKarp(s, substr string) int {
	n := len(substr)
	hashss, pow := hashStr(substr)

	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}

	if hashss == h && s[:n] == substr {
		return 0
	}

	i := n
	for i < len(s) {
		h = h*primeRK + uint32(s[i])
		h -= pow*uint32(s[i-n])
		i++

		if hashss == h && s[i-n:i] == substr {
			return i-n
		}
	}

	return -1
}


func strStr0(haystack string, needle string) int {
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
