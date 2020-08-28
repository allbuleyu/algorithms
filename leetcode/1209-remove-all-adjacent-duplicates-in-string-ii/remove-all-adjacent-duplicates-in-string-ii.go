package prob1209

func removeDuplicates(s string, k int) string {
	return helpTwoPointers(s, k)
}

func helpTwoPointers(s string, k int) string {
	b := []byte(s)
	stack := make([]int, 0)
	j := 0
	for i := 0; i < len(s); i++ {
		b[j] = b[i]

		if j > 0 && b[j] == b[j-1] {
			times := stack[len(stack)-1]
			if times + 1 == k {
				j -= k
				stack = stack[:len(stack)-1:cap(stack)]
			}else {
				stack[len(stack)-1]++
			}
		}else {
			stack = append(stack, 1)
		}

		j++
	}

	return string(b[:j])
}

func helpStack(s string, k int) string {
	stack := make([]byte, 0)
	dupStart := 0
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			dupEnd := len(stack)
			if dupEnd-dupStart+1 == k {
				stack = stack[:dupStart:cap(stack)]

				if len(stack) == 0 {
					dupStart = 0
				}else {
					for dupStart = len(stack)-1; dupStart > 0; dupStart-- {
						if stack[dupStart] != stack[dupStart-1] {
							break
						}
					}
				}

				continue
			}
		}else {
			dupStart = len(stack)
		}

		stack = append(stack, s[i])
	}

	return string(stack)
}