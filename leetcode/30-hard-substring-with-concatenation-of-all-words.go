package leetcode

//https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/
//You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.
//
//For example, given:
//s: "barfoothefoobarman"
//words: ["foo", "bar"]
//
//You should return the indices: [0,9].
//(order does not matter).

func findSubstring(s string, words []string) []int {
	m := make(map[string]int)
	count := 0
	wordLen := 0
	for _, word := range words {
		if wordLen == 0 {
			wordLen = len(word)
		}

		m[word]++
		count++
	}

	strLen := len(words) * wordLen
	res := make([]int, 0)
	start := 0
	for i:=wordLen; i < len(s);i++ {
		m[s[i-wordLen:i]]--
		if m[s[i-wordLen:i]] >= 0 {
			count--
		}

		if count == 0 {
			res = append(res, start)
		}

		if i+1-strLen == start {
			m[s[start:start+wordLen]]++
			if m[s[start:start+wordLen]] > 0 {
				count++
			}
			start++
		}
	}

	return res
}

func findSubstring1(s string, words []string) []int {
	hs := [58]int{}
	count := 0
	for _, word := range words {
		for _, w := range word {
			hs[w-'A']++
			count++
		}
	}
	sbuLen := len(words[0])
	strLen := count
	start := 0
	res := make([]int, 0)
	for i, v := range s {
		hs[v-'A']--
		if hs[v-'A'] >= 0 {
			count--
		}

		if count == 0 {
			res = append(res, start)

			for j:= start; j < start+sbuLen; j++ {
				hs[s[j]-'A']++
			}

			start += sbuLen
			count += sbuLen
			i += sbuLen
			continue
		}

		if i+1-strLen == start {
			hs[s[start]-'A']++
			if hs[s[start]-'A'] > 0 {
				count++
			}

			start++
		}
	}

	return res
}

func FindSubstring(s string, words []string) []int {
	return findSubstring(s , words)
}