package prob0459

import "strings"

func repeatedSubstringPattern(s string) bool {
	n := len(s)

	if n <= 1 {
		return false
	}

	subStrLen := 1

	// 时间复杂度
	// first n, second n/2, third n/3 ... = 2n = O(n)
	GS:
	for subStrLen <= n/2 {
		if n % subStrLen != 0 {
			subStrLen++
			continue GS
		}

		subStr := s[0:subStrLen]
		for i := 0; i + subStrLen <= n; {
			if s[i:i+subStrLen] != subStr {
				subStrLen++
				goto GS
			}
			i += subStrLen
		}

		return true
	}

	return false
}


// 至少包含一组相同的子字符串
func srepeatedSubstringPattern1(s string) bool {
	if len(s) == 0 {
		return false
	}

	size := len(s)

	ss := (s + s)[1 : size*2-1]

	return strings.Contains(ss, s)
}