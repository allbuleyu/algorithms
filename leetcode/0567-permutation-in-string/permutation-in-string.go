package prob0567

//https://leetcode.com/problems/permutation-in-string/description/
//Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.
//Example 1:
//Input:s1 = "ab" s2 = "eidbaooo"
//Output:True
//Explanation: s2 contains one permutation of s1 ("ba").
//Example 2:
//Input:s1= "ab" s2 = "eidboaoo"
//Output: False
//Note:
//The input strings only contain lower case letters.
//The length of both given strings is in range [1, 10,000].
//next challenge 76
func checkInclusion(s1 string, s2 string) bool {
	return slidingWindow(s1, s2)
}

func CheckInclusion(s1 string, s2 string) bool {
	var hs [26]int
	var count, start int
	for _, v := range s1 {
		hs[v-'a']++
		count++
	}

	for i, v := range s2 {
		hs[v-'a']--
		if hs[v-'a'] >=0 {
			count--
		}

		if i - start >= len(s1) {
			hs[s2[start]-'a']++
			if hs[s2[start]-'a'] > 0 {
				count++
			}
			start++
		}

		if count == 0 {
			return true
		}
	}

	return false
}

func slidingWindow(s1, s2 string) (ans bool) {
	m, n := len(s1), len(s2)
	if m > n {
		return
	}
	hs := [26]int{}
	for i := 0; i < m; i++ {
		hs[s1[i]-'a']++
	}
	count := m

	for i, j := 0, 0; i < n; i++ {
		hs[s2[i]-'a']--
		if hs[s2[i]-'a'] >= 0 {
			count--
		}

		if count == 0 {
			ans = true
			break
		}

		// sliding the window
		if i-j+1 == m {
			hs[s2[j]-'a']++
			if hs[s2[j]-'a'] > 0 {
				count++
			}

			j++
		}

	}

	return
}