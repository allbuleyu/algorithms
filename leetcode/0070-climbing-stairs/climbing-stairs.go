package prob0070


//1 1 -> 1
//2 (11)*2, 2 -> 2
//3 (111) * 3, 12, 21  -> 3
//4 (1111) * 4, (112) * 2, (121) * 2, (211)*2, 22 -> 5
//5 (11111) * 5, (1112) * 3, 1121 * 3, 1211 * 3, 2111 * 3, 122, 212, 221 -> 8

// 推导过程
// 设s为解的集合, s[i] 为i个台阶的解
// 求s[i+1] 的解
// 因为只有两种走法, 1步或者2步, 因此 i+1的走法可以看成 在走i步的基础上走了1步, 在i-1的基础上走了两步,因为没有更多的走法
// 所以s[i+1] = s[i] + s[i-1]
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

// 优化,把空间复杂度从O(n) -> O(1)
func climbStairsOptimize(n int) int {
	if n == 1 {
		return 1
	}

	var step, step1, step2 int

	step1 = 1
	step2 = 0
	for i := 0; i < n; i++ {
		step = step1 + step2

		step1, step2 = step, step1
	}

	return step
}