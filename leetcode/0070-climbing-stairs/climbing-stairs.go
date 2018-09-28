package prob0070


//1 1 -> 1
//2 (11)*2, 2 -> 2
//3 (111) * 3, 12, 21  -> 3
//4 (1111) * 4, (112) * 2, (121) * 2, (211)*2, 22 -> 5
//5 (11111) * 5, (1112) * 3, 1121 * 3, 1211 * 3, 2111 * 3, 122, 212, 221 -> 8

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	steps := make([]int, n)
	steps[0] = 1
	steps[1] = 2

	for i := 2; i < n; i++ {
		steps[i] = steps[i-1] + steps[i-2]
	}

	return steps[n-1]
}