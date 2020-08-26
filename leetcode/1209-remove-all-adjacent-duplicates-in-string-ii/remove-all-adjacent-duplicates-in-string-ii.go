package prob1209

func removeDuplicates(s string, k int) string {
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