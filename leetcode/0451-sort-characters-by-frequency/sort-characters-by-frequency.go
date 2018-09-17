package prob0451

import "sort"

//https://leetcode.com/problems/sort-characters-by-frequency/description/
//Given a string, sort it in decreasing order based on the frequency of characters.
//
//Example 1:
//
//Input:
//"tree"
//
//Output:
//"eert"
//
//Explanation:
//'e' appears twice while 'r' and 't' both appear once.
//So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
//Example 2:
//
//Input:
//"cccaaa"
//
//Output:
//"cccaaa"
//
//Explanation:
//Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
//Note that "cacaca" is incorrect, as the same characters must be together.
//Example 3:
//
//Input:
//"Aabb"
//
//Output:
//"bbAa"
//
//Explanation:
//"bbaA" is also a valid answer, but "Aabb" is incorrect.
//Note that 'A' and 'a' are treated as two different characters.
//Next chanllenge 347
func frequencySort(s string) string {
	r := ['z' + 1]int{}
	for i := range s {
		r[s[i]]++
	}

	segs := make([]string, 0, len(s))
	for i := range r {
		if r[i] == 0 {
			continue
		}
		segs = append(segs, makeString(byte(i), r[i]))	//把重复次数的字符串放进数组
	}

	sort.Sort(segments(segs))	//

	res := ""
	for _, s := range segs {
		res += s
	}

	return res
}

// 返回由 n 个 b 组成的字符串
func makeString(b byte, n int) string {
	bytes := make([]byte, n)
	for i := range bytes {
		bytes[i] = b
	}
	return string(bytes)
}

// segments 实现了 sort.Interface 接口
type segments []string

func (s segments) Len() int { return len(s) }

func (s segments) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return s[i] < s[j]
	}
	return len(s[i]) > len(s[j])
}

func (s segments) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func FrequencySort(s string) string {
	return frequencySort(s)
}