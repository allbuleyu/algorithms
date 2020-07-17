package prob0557

func reverseWords(s string) string {
	sb := make([]byte, len(s))
	ans := make([]byte, 0, len(s))
	for i, j := 0, len(s)-1; i<=j; i,j = i+1,j-1 {
		sb[i], sb[j] = s[j], s[i]
	}

	for i := len(s)-1; i >= 0; i-- {
		j := i
		for ; j >= 0; j-- {
			if j == 0 {
				ans = append(ans, sb[j:i+1]...)
				break
			}

			if sb[j] == ' ' {
				ans = append(ans, sb[j+1:i+1]...)
				ans = append(ans, ' ')
				break
			}
		}
		i=j
	}

	return string(ans)
}
