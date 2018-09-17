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
	n1, n2 := len(s1), len(s2)
	dp := make([]int, n2+1)

	for j := n2 -1; j >= 0; j-- {
		dp[j] = dp[j+1] + int(s2[j])
	}


	for i := n1 - 1; i >= 0; i-- {
		tmp := make([]int, n2+1)
		if i == n1 - 1 {
			for j := n2 -1; j >= 0; j-- {
				tmp[j] = tmp[j+1] + int(s2[j])
			}
		}

		tmp[n2] = int(s1[i])
		for j := n2 -1; j >= 0; j-- {
			if s1[i] == s2[j] {
				tmp[j] = dp[j+1]
			}else {
				tmp[j] = min(dp[j]+int(s1[i]), tmp[j+1]+int(s2[j]))
			}
		}
		dp = tmp
	}

	return dp[0]
}

func minimumDeleteSum1(s1 string, s2 string) int {
	n1, n2 := len(s1), len(s2)
	dp := make([]string, n2+1)

	for j := n2 -1; j >= 0; j-- {
		dp[j] = dp[j+1] + string(s2[j])
	}


	for i := n1 - 1; i >= 0; i-- {
		tmp := make([]string, n2+1)
		if i == n1 - 1 {
			for j := n2 -1; j >= 0; j-- {
				tmp[j] = tmp[j+1] + string(s2[j])
			}
		}

		tmp[n2] = string(s1[i])
		for j := n2 -1; j >= 0; j-- {
			if s1[i] == s2[j] {
				tmp[j] = dp[j+1]
			}else {
				tmp[j] = min1(dp[j]+string(s1[i]), tmp[j+1]+string(s2[j]))
			}
		}
		dp = tmp
	}

	return 1
}

func MinimumDeleteSum(s1 string, s2 string) int {
	return minimumDeleteSum1(s1 , s2 )
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min1(a, b string) string {
	var sum1, sum2 int
	for i := range a {
		sum1 += int(a[i])
	}

	for i := range b {
		sum2 += int(b[i])
	}

	if sum1 > sum2 {
		return b
	}

	return a
}