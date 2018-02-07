package leetcode

//Given a non-empty array of integers, return the k most frequent elements.
//
//For example,
//Given [1,1,1,2,2,3] and k = 2, return [1,2].
//
//Note:
//You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
//Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}

	fr := make([][]int, len(m)+1)

	for i, v := range m {
		fr[v] = append(fr[v], i)
	}

	r := make([]int, 0, k)
	for i := len(m); i >= 0; i-- {
		if len(fr[i]) == 0 {
			continue
		}

		r = append(r, fr[i]...)
		if len(r) >= k {
			return r[:k]
		}
	}

	return r
}

func TopKFrequent(nums []int, k int) []int {
	return topKFrequent(nums, k)
}