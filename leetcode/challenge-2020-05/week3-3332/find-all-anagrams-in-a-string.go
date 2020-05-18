package week3_3332

//Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.
//
//Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.
//
//The order of output does not matter.
//
//Example 1:
//
//Input:
//s: "cbaebabacd" p: "abc"
//
//Output:
//[0, 6]
//
//Explanation:
//The substring with start index = 0 is "cba", which is an anagram of "abc".
//The substring with start index = 6 is "bac", which is an anagram of "abc".
//Example 2:
//
//Input:
//s: "abab" p: "ab"
//
//Output:
//[0, 1, 2]
//
//Explanation:
//The substring with start index = 0 is "ab", which is an anagram of "ab".
//The substring with start index = 1 is "ba", which is an anagram of "ab".
//The substring with start index = 2 is "ab", which is an anagram of "ab".

func findAnagrams(s string, p string) []int {
	var hs [26]int
	res := make([]int, 0)
	for _, v := range p {
		hs[v-'a']++
	}

	var start, count int
	count = len(p)
	for i := 0; i < len(s); i++ {
		v := s[i]
		hs[v-'a']--
		if hs[v-'a'] >= 0 {
			count--
		}

		if count == 0 {
			res = append(res, start)
		}

		if i-start+1 == len(p) {
			hs[s[start]-'a']++
			if hs[s[start]-'a'] > 0 {
				count++
			}

			start++
		}
	}

	return res
}
