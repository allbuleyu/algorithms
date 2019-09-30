package prob1014

func maxScoreSightseeingPair(A []int) int {
	res := 0
	cur := 0

	// a[i]+a[j] + i - j = a[i] + i + a[j] - j
	for _, v := range A {
		res = max(res, cur+v)
		cur = max(cur, v) - 1
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}