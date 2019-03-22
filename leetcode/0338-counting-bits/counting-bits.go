package prob0338

const (
	m1 = 0x55555555
	m2 = 0x33333333
	m4 = 0x0f0f0f0f
	m8 = 0x00ff00ff
	m16 = 0x0000ffff
)

func countBits(num int) []int {
	return solution2(num)
}

func solution2(num int) []int {
	dp := make([]int, num+1)

	for i:= 1; i <= num; i++ {
		dp[i] = dp[i>>1] + (i & 1)
	}

	return dp
}

func solution1(num int) []int {
	res := make([]int, num+1)

	for i:= 0; i <= num; i++ {
		res[i] = hammingWeight(i)
	}

	return res
}

// bench 100000000	        3.5 ns/op
// https://en.wikipedia.org/wiki/Hamming_weight
func hammingWeight(num int) int {
	num = (num & m1) + (num >> 1) & m1		// 计算2位上1的个数
	num = (num & m2) + (num >> 2) & m2		// 计算4位上1的个数
	num = (num & m4) + (num >> 4) & m4		// 计算8位上1的个数
	num = (num & m8) + (num >> 8) & m8		// 计算16位上1的个数
	num = (num & m16) + (num >> 16) & m16		// 计算32位上1的个数

	return int(num)
}