package leetcode

import "sort"

//https://leetcode.com/problems/top-k-frequent-words/description/
//Given a non-empty list of words, return the k most frequent elements.
//
//Your answer should be sorted by frequency from highest to lowest. If two words have the same frequency, then the word with the lower alphabetical order comes first.
//
//Example 1:
//Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
//Output: ["i", "love"]
//Explanation: "i" and "love" are the two most frequent words.
//Note that "i" comes before "love" due to a lower alphabetical order.
//Example 2:
//Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
//Output: ["the", "is", "sunny", "day"]
//Explanation: "the", "is", "sunny" and "day" are the four most frequent words,
//with the number of occurrence being 4, 3, 2 and 1 respectively.
//Note:
//You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
//Input words contain only lowercase letters.
//Follow up:
//Try to solve it in O(n log k) time and O(n) extra space.
// next challenge 438, 508
func topKFrequentStr(words []string, k int) []string {
	m := make(map[string]int)
	for i := range words {
		m[words[i]]++
	}

	fr := make([][]string, len(words) + 1)
	for i, v := range m {
		fr[v] = append(fr[v], i)
	}

	r := make([]string, 0, k)
	for i := len(fr)-1; i >= 0; i-- {
		sort.Sort(wordStrings(fr[i]))
		r = append(r, fr[i]...)
		if len(r) >= k {
			return r[:k]
		}
	}

	return r
}

func TopKFrequentStr(words []string, k int) []string {
	return topKFrequentStr(words, k)
}

type wordStrings []string

func (s wordStrings) Len() int { return len(s) }

func (s wordStrings) Less(i, j int) bool {
	var l int
	if len(s[i]) > len(s[j]) {
		l = len(s[j])
	}else {
		l = len(s[i])
	}

	for k := 0; k < l; k++ {
		if s[i][k] < s[j][k] {
			return true
		}else if s[i][k] > s[j][k] {
			return false
		}
	}

	return len(s[i]) < len(s[j])
}

func (s wordStrings) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

