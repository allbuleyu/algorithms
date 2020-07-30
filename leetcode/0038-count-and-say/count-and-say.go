package prob0038

import "fmt"

func countAndSay(n int) string {
	if n < 1 {
		return ""
	}

	res := "1"
	for n > 1 {
		tmp := ""
		for i := 0; i < len(res); i++ {
			times := 1
			for j := i+1; j < len(res) && res[i] == res[j]; j++ {
				times++
				i++
			}
			tmp += fmt.Sprintf("%d%s", times, string(res[i]))
		}
		res = tmp
		n--
	}

	return res
}