package prob0242

import (
	"sort"
	"strings"
)

//https://leetcode.com/problems/valid-anagram/description/
//Given two strings s and t, write a function to determine if t is an anagram of s.
//
//For example,
//s = "anagram", t = "nagaram", return true.
//s = "rat", t = "car", return false.
//
//Note:
//You may assume the string contains only lowercase alphabets.
//
//Follow up:
//What if the inputs contain unicode characters? How would you adapt your solution to such case?
//next challenge 49
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return  false
	}

	var hs [26]int
	for _, v := range t {
		hs[v-'a']++
	}

	for _, v := range s {
		hs[v-'a']--
		if hs[v-'a'] < 0 {
			return false
		}
	}

	return true
}

// hash table is fastest
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return  false
	}

	var hs [26]int
	for _, v := range t {
		hs[v-'a']++
	}

	for _, v := range s {
		hs[v-'a']--
		if hs[v-'a'] < 0 {
			return false
		}
	}

	return true
}

// sorted string
func isAnagram(s string, t string) bool {
	sliceS := strings.Split(s, "")
	sliceT := strings.Split(t, "")
	sort.Strings(sliceS)
	sort.Strings(sliceT)
	if strings.Join(sliceS, "") != strings.Join(sliceT, "") {
		return false
	}

	return true
}

func IsAnagram(s string, t string) bool {
	return isAnagram(s, t)
}