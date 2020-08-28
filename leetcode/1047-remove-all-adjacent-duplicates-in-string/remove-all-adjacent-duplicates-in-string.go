package prob1047

func removeDuplicates(s string) string {
	return helpTwoPointers(s)
}

func helpTwoPointers(s string) string {
	j := 0
	b := []byte(s)
	for i := 0; i < len(s); i++ {
		b[j] = b[i]

		if j > 0 && b[j] == b[j-1] {
			j -= 2
		}

		j++
	}

	return string(b[:j])
}

func helpStack(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1:cap(stack)]
		}else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}