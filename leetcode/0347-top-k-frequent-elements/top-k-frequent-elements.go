package prob0347

//https://leetcode.com/problems/top-k-frequent-elements/
//Given a non-empty array of integers, return the k most frequent elements.
//
//For example,
//Given [1,1,1,2,2,3] and k = 2, return [1,2].
//
//Note:
//You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
//Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
// next challenge 192(bash), 215,659,692
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}

	// len(nums) + 1 解释:  如果一个数组为只有1个元素,或者都是相等的元素,那么fr[v] = fr[len(nums)]
	// input [1,1], m => m[1]=2, fr[2] = 1 数组下标2 = len([1,1])+1 所以fr得长度为3
	fr := make([][]int, len(nums)+1)

	for i, v := range m {
		fr[v] = append(fr[v], i)
	}

	r := make([]int, 0, k)
	for i := len(fr)-1; i >= 0; i-- {
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