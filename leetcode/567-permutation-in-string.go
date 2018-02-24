package leetcode

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

func CheckInclusion(s1 string, s2 string) bool {
	return checkInclusion(s1, s2)
}