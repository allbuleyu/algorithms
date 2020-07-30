package prob0443

import "strconv"

func compress(chars []byte) int {
	ans := 0
	for i := 0; i < len(chars); i++ {
		times := 1
		for j := i+1; j < len(chars) && chars[i] == chars[j]; j++ {
			times++
			i++
		}

		chars[ans] = chars[i]
		ans++

		if times > 1 {
			t := strconv.Itoa(times)
			for k := 0; k < len(t); k++ {
				chars[ans] = t[k]
				ans++
			}
		}
	}

	return ans
}