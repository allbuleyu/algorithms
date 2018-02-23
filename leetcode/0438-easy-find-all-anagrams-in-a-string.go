package leetcode

//https://leetcode.com/problems/find-all-anagrams-in-a-string/description/
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
//difficult easy
func findAnagrams(s string, p string) []int {
	var r, rc [123]int
	for _, v := range p {
		r[byte(v)]++
	}
	copy(r, rc)
	numOfDeference := len(p)

	res := make([]int, 0)
	for left, right := 0,0; right < len(s); right++ {
		v := s[right]

		if r[byte(v)] == 0 {
			left++
			numOfDeference = len(p)
			r = rc
			continue
		}

		r[byte(v)]--
		if s[byte(v)] >= 0 {
			numOfDeference--
		}

		if numOfDeference == 0 {
			res = append(res, left)
			right = left
		}


	}

	return res
}

func FindAnagrams(s string, p string) []int {
	return findAnagrams(s, p)
}