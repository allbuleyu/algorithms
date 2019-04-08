package prob0350

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	return iteration(nums1, nums2)
}

func iteration(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	res := make([]int, 0)
	start := 0
	for i := 0; i < len(nums1); i++ {

		for j := start; j < len(nums2);j++ {
			if nums1[i] == nums2[j] {
				res = append(res, nums1[i])
				start = j+1
				break
			}

		}
	}

	return res
}