package prob0712

//https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/description/
//Given two strings s1, s2, find the lowest ASCII sum of deleted characters to make two strings equal.
//
//Example 1:
//Input: s1 = "sea", s2 = "eat"
//Output: 231
//Explanation: Deleting "s" from "sea" adds the ASCII value of "s" (115) to the sum.
//Deleting "t" from "eat" adds 116 to the sum.
//At the end, both strings are equal, and 115 + 116 = 231 is the minimum sum possible to achieve this.
//Example 2:
//Input: s1 = "delete", s2 = "leet"
//Output: 403
//Explanation: Deleting "dee" from "delete" to turn the string into "let",
//adds 100[d]+101[e]+101[e] to the sum.  Deleting "e" from "leet" adds 101[e] to the sum.
//At the end, both strings are equal to "let", and the answer is 100+101+101+101 = 403.
//If instead we turned both strings into "lee" or "eet", we would get answers of 433 or 417, which are higher.
//Note:
//
//0 < s1.length, s2.length <= 1000.
//All elements of each string will have an ASCII value in [97, 122].

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		if i > 0 {
			dp[i][0] = int(s1[i-1]) + dp[i-1][0]
		}
	}

	for i := 1; i < n + 1;i++ {
		dp[0][i] = int(s2[i-1]) + dp[0][i-1]
	}



	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else {
				dp[i][j] = min(dp[i][j-1]+int(s2[j-1]), dp[i-1][j]+int(s1[i-1]))
			}
		}
	}


	return dp[m][n]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}