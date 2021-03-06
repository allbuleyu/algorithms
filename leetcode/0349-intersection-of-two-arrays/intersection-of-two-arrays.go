package prob0349

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	return iteration(nums1, nums2)
}

func iteration(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	res := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		if i > 0 && nums1[i] == nums1[i-1] {
			continue
		}

		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				res = append(res, nums1[i])
				break
			}

		}
	}

	return res
}

// solution with bs
func binarySearch()  {

}