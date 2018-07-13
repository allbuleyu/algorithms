package ch15

// 动态规划,斐波那契 Fibonacci
// 15.1.5
func DynamicFibo(n int) int64 {
	r := make([]int64, n+1)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	r[0] = 0
	r[1] = 1

	for j := 2; j <= n; j++ {
		var s int64
		s = r[j-1] + r[j-2]

		r[j] = s
	}

	return r[n]
}
