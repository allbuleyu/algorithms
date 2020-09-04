package prob0394

import "strconv"

func decodeString(s string) string {
	return helpStack(s)
}

func helpStack(s string) string {
	stackTimes := make([]int, 0)
	sb := make([][]byte, 0)

	b := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		j := i
		for s[j] >= '0' && s[j] <= '9' {
			j++
		}

		// encounting nums
		if j > i {
			times, _ := strconv.Atoi(s[i:j])
			stackTimes = append(stackTimes, times)
			i = j -1
			continue
		}

		if s[i] == '[' {
			sb = append(sb, []byte{})
		}else if s[i] == ']' {
			if len(sb) == 1 {
				b = append(b, generateStr(sb[0], stackTimes[0])...)
				stackTimes = stackTimes[:0:cap(stackTimes)]
				sb = sb[:0:cap(sb)]
			}else {
				last := len(sb)-1
				sb[last-1] = append(sb[last-1], generateStr(sb[last], stackTimes[last])...)
				stackTimes = stackTimes[:last:cap(stackTimes)]
				sb = sb[:last:cap(sb)]
			}
		}else {
			if len(sb) == 0 {
				b = append(b, s[i])
			}else {
				sb[len(sb)-1] = append(sb[len(sb)-1], s[i])
			}

		}

	}

	return string(b)
}

func generateStr(b []byte, n int) []byte {
	ans := make([]byte, 0, len(b)*n)
	for i := 0; i < n; i++ {
		ans = append(ans, b...)
	}

	return ans
}