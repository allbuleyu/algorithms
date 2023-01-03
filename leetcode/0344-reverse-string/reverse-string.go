package prob0344

func reverseString(s []byte) {
	helpDivideAndConquer(s)
}

func helpDivideAndConquer(s []byte) {
	if len(s) < 2 {
		return
	}

	divideAndConquer(s, 0, len(s)-1)
}

// abcde
// abc de
// ab c d e
func divideAndConquer(s []byte, start, end int) {
	if start == end {
		return
	}

	d := (start + end) / 2
	divideAndConquer(s, start, d)
	divideAndConquer(s, d+1, end)

	for i, j := start, end; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
