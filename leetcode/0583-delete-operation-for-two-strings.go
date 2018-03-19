package leetcode

import "math"

//https://leetcode.com/problems/delete-operation-for-two-strings/description/
//Given two words word1 and word2, find the minimum number of steps required to make word1 and word2 the same, where in each step you can delete one character in either string.
//
//Example 1:
//Input: "sea", "eat"
//Output: 2
//Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
//Note:
//The length of given words won't exceed 500.
//Characters in given words can only be lower-case letters.

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	return m + n - 2 * lcs(word1, word2, m, n)
}


func lcs(word1, word2 string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}else if word1[m-1] == word2[n-1] {
		return 1 + lcs(word1, word2, m-1, n-1)
	}else {
		max := math.Max(float64(lcs(word1, word2, m-1, n)), float64(lcs(word1, word2, m, n-1)))
		return int(max)
	}
}


// 求出最大公共子串的长度
func minDistance1(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m;i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}

	return m + n - 2 * dp[m][n]
}

// 直接算删除的步骤
// 时间复杂度 O(m*n) 空间复杂度 O(m*n)
func minDistance2(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m;i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = i+j
			}else if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + 1
			}
		}
	}

	return dp[m][n]
}

// 直接算删除的步骤 2的进阶
// 时间复杂度 O(m*n) 空间复杂度 (n)
func minDistance3(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([]int, n+1)


	for i := 0; i <= m;i++ {
		tmp := make([]int, n+1)
		for j := 0; j <= n; j++ {

			if i == 0 || j == 0 {
				tmp[j] = i+j
			}else if word1[i-1] == word2[j-1] {
				tmp[j] = dp[j-1]
			}else {
				tmp[j] = int(math.Min(float64(tmp[j-1]), float64(dp[j]))) + 1
			}
		}

		dp = tmp
	}

	return dp[n]
}

func MinDistance(word1 string, word2 string) int {
	return minDistance3(word1, word2)
}