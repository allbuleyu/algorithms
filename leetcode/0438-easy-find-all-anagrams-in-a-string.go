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
//difficult easy 这个难度并不简单,应该算中等难度 567比这个还要简单些
//next challenge 242,567
func check(a,b,c [26]int) bool{
	for i:=0;i<26;i++{
		if a[i]-b[i] != c[i]{
			return false
		}
	}
	return true
}

func findAnagrams1(s string, p string) []int {
	if len(s)<len(p){
		return []int{}
	}
	ret := []int{}
	start, end, mark := [26]int{},[26]int{},[26]int{}

	var ck func(a,b,c [26]int) bool
	ck = func(end, start, mark [26]int) bool {
		for i:=0;i<26;i++{
			if end[i]-start[i] != mark[i]{
				return false
			}
		}
		return true
	}

	for i:=0;i<len(p);i++{
		mark[p[i]-'a'] ++
	}
	for i:=0;i<len(p);i++{
		end[s[i]-'a'] ++
	}
	if ck(end, start, mark){
		ret = append(ret,0)
	}
	for i:=len(p);i<len(s);i++{
		start[s[i-len(p)]-'a'] ++
		end[s[i]-'a']++
		if ck(end,start,mark){
			ret = append(ret,i-len(p)+1)
		}
	}
	return ret
}
// 移动窗口
func findAnagrams(s string, p string) []int {
	hs := [256]int{}
	start := 0
	count := 0
	res := make([]int, 0)

	for _, v := range p {
		hs[v]++
		count++
	}

	for i, v := range s {
		hs[v]--
		if hs[v] >= 0 {
			count--
		}

		if i - start >= len(p) {
			hs[s[start]]++
			if hs[s[start]] > 0 {
				count++
			}
			start++
		}

		if count == 0 {
			res = append(res, start)
		}
	}

	return res
}

func FindAnagrams(s string, p string) []int {
	return findAnagrams(s, p)
}