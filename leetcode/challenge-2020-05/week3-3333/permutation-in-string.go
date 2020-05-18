package week3_3333

//Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.
//
// 
//
//Example 1:
//
//Input: s1 = "ab" s2 = "eidbaooo"
//Output: True
//Explanation: s2 contains one permutation of s1 ("ba").
//Example 2:
//
//Input:s1= "ab" s2 = "eidboaoo"
//Output: False
// 
//
//Note:
//
//The input strings only contain lower case letters.
//The length of both given strings is in range [1, 10,000].
func checkInclusion(s1 string, s2 string) bool {
	return helpSlideWindow(s1, s2)
}

func helpSlideWindow(s1 string, s2 string) bool {
	hs := [26]int{}
	for _, v := range s1 {
		hs[v-'a']++
	}
	
	count := len(s1)
	start := 0
	for i, v := range s2 {
		hs[v-'a']--
		if hs[v-'a'] >= 0 {
			count--
		}
		
		if count == 0 {
			return true
		}
		
		if i-start+1 == len(s1) {
			hs[s2[start]-'a']++
			if hs[s2[start]-'a'] > 0 {
				count++
			}
			
			start++
		}
	}

	return false
}
