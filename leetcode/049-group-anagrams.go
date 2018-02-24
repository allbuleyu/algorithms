package leetcode

import (
	"strings"
	"sort"
)

//https://leetcode.com/problems/group-anagrams/description/
//Given an array of strings, group anagrams together.
//
//For example, given: ["eat", "tea", "tan", "ate", "nat", "bat"],
//Return:
//
//[
//["ate", "eat","tea"],
//["nat","tan"],
//["bat"]
//]
//Note: All inputs will be in lower-case.

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	var ans [][]string
	for _, v := range strs {
		sv := strings.Split(v, "")
		sort.Strings(sv)
		key := strings.Join(sv, "")

		m[key] = append(m[key], v)
	}

	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
}

func GroupAnagrams(strs []string) [][]string {
	return groupAnagrams(strs)
}