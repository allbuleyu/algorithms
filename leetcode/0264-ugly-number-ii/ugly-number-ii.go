package prob0264


func nthUglyNumber(n int) int {
	return dynamicProgramming(n)
}

func dynamicProgramming(n int) int {
	res := make([]int, n)
	res[0] = 1

	var u2, u3, u5 int
	for i := 1; i < n; i++ {
		res[i] = min(res[u2]*2, min(res[u3] * 3, res[u5] *5))

		if res[i] == res[u2] * 2 {
			u2++
		}

		if res[i] == res[u3] * 3 {
			u3++
		}

		if res[i] == res[u5] * 5 {
			u5++
		}
	}

	return res[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}