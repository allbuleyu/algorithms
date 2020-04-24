package _528_week_01_3288

import (
	"bytes"
	"sort"
	"strconv"
)

//https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/528/week-1/3288/
//Given an array of strings, group anagrams together.
//
//Example:
//
//Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
//Output:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//Note:
//
//All inputs will be in lowercase.
//The order of your output does not matter.
func groupAnagrams(strs []string) [][]string {
	return groupAnagramsOptimize(strs)
}

type ASCIIMap [26]int
func (m ASCIIMap) toKey() string {
	b := bytes.Buffer{}
	for _, v := range m {
		b.WriteString(strconv.Itoa(v))
		b.WriteByte('#')
	}

	return b.String()
}

func strToKey(str string) string {
	m := ASCIIMap{}

	for _, v := range str {
		m[v-'a']++
	}

	return m.toKey()
}

func groupAnagramsOptimize(strs []string) [][]string {
	hs := make(map[string][]string)
	for _, str := range strs {
		key := strToKey(str)
		if _, ok := hs[key]; !ok {
			hs[key] = make([]string, 0)
		}

		hs[key] = append(hs[key], str)
	}

	ans := make([][]string, 0)
	for _, v := range hs {
		ans = append(ans , v)
	}

	return ans
}

func g1(strs []string) [][]string {
	hs := make(map[string][]string)

	for _, str := range strs {
		tb := []byte(str)
		sort.Slice(tb, func(i, j int) bool {
			if tb[i] < tb[j] {
				return true
			}

			return false
		})

		if _, ok := hs[string(tb)];!ok {
			hs[string(tb)] = make([]string, 0)
		}

		hs[string(tb)] = append(hs[string(tb)], str)
	}

	ans := make([][]string, 0)
	for _, args := range hs {
		ans = append(ans, args)
	}

	return ans
}
