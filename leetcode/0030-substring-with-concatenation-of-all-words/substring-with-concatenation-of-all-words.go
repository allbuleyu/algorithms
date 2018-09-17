package prob0030

//https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/
//You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.
//
//For example, given:
//s: "barfoothefoobarman"
//words: ["foo", "bar"]
//
//You should return the indices: [0,9].
//(order does not matter).
// 移动窗口,把words里面的单词当做一个字母的思路来解
func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	hashMap := make(map[string]int, len(words))

	lens, lenWords, subStrLen := len(s), len(words), len(words[0])

	count := lenWords
	initFunc := func() {
		for _, word := range words {
			hashMap[word] = 0
		}

		for _, word := range words {
			hashMap[word]++
		}

		count = lenWords
	}

	for i := 0; i < subStrLen; i++ {
		index := i

		moveIndex := func() {
			hashMap[s[index:index+subStrLen]]++
			index += subStrLen
			count++
		}

		initFunc()

		for index + lenWords * subStrLen <= lens {
			start := index + (lenWords - count) * subStrLen
			end := index + (lenWords - count + 1) * subStrLen
			word := s[start:end]
			numWord, ok := hashMap[word]

			if !ok {
				index += subStrLen

				if count != lenWords {
					initFunc()
				}
			}else if numWord == 0 {
				//hashMap[word]++
				//count++
				//index += subStrLen
				moveIndex()
			}else {
				hashMap[word]--
				count--

				if count == 0 {
					res = append(res, index)

					//hashMap[word]++
					//index += subStrLen
					//count++
					moveIndex()
				}
			}
		}
	}

	return res
}

func FindSubstring(s string, words []string) []int {
	return findSubstring(s , words)
}