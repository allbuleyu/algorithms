package prob0718

//
//https://leetcode.com/problems/maximum-length-of-repeated-subarray/description/
//Given two integer arrays A and B, return the maximum length of an subarray that appears in both arrays.
//
//Example 1:
//Input:
//A: [1,2,3,2,1]
//B: [3,2,1,4,7]
//Output: 3
//Explanation:
//The repeated subarray with maximum length is [3, 2, 1].
//Note:
//1 <= len(A), len(B) <= 1000
//0 <= A[i], B[i] < 100

func findLength(A []int, B []int) int {
	return bottomUp(A, B)
}

func bottomUp(a, b []int) int {
	res := 0
	dp := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		dp[i] = make([]int, len(b)+1)
	}

	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {
			if a[i] == b[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				res = max(res, dp[i][j])
			}
		}
	}

	return res
}

var memory [][]int
var ans int

func helperTopDown(a, b []int) int {
	memory = make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		memory[i] = make([]int, len(b))

		for j := 0; j < len(b); j++ {
			memory[i][j] = -1
		}
	}
	ans = 0
	topDown(a, b, 0, 0)

	return ans
}

func topDown(a, b []int, i, j int) int {
	if i == len(a) || j == len(b) {
		return 0
	}

	if memory[i][j] != -1 {
		return memory[i][j]
	}

	var res int
	if a[i] == b[j] {
		res = topDown(a, b, i+1, j+1) + 1
	} else {
		res = max(topDown(a, b, i+1, j), topDown(a, b, i, j+1))
	}
	ans = max(ans, res)
	memory[i][j] = res

	return memory[i][j]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
