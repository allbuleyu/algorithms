package prob0451

import "sort"

func frequencySort(s string) string {
	return helpBucketSort(s)
}

func helpBucketSort(s string) string {
	hs := [128]int{}
	maxLen := 0
	for i := 0; i < len(s); i++ {
		v := hs[s[i]]
		hs[s[i]]++

		maxLen = max(maxLen, v+1)
	}

	bucket := make(map[int][]byte)
	for i := 0; i < 127; i++ {
		if hs[i] == 0 {
			continue
		}

		if _, ok := bucket[hs[i]]; !ok {
			bucket[hs[i]] = make([]byte, 0)
		}

		bucket[hs[i]] = append(bucket[hs[i]], byte(i))
	}

	bb := make([]byte, 0, len(s))
	for i := maxLen; i > 0; i-- {
		if bs, ok := bucket[i]; ok {
			for j := 0; j < len(bs); j++ {
				bb = append(bb, generateString(bs[j], i)...)
			}
		}
	}

	return string(bb)
}

func generateString(b byte, t int) []byte {
	bb := make([]byte, t)
	for i := 0; i < t; i++ {
		bb[i] = b
	}

	return bb
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// hash map + sort
func frequencySort1(s string) string {
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