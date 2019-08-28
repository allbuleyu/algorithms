package prob0350

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	return forward1(nums1, nums2)
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

// Follow up
// What if the given array is already sorted? How would you optimize your algorithm?
// What if nums1's size is small compared to nums2's size? Which algorithm is better?
//
func forward1(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)


	ans := make([]int, 0)
	m, n := len(nums1)-1, len(nums2)-1

	// O(N) = O(min(m/n))
	for m >= 0 && n >= 0 {
		// skip duplicate element
		//for m > 0 && nums1[m] == nums1[m-1] {
		//	m--
		//}
		//
		//for n > 0 && nums2[n] == nums2[n-1] {
		//	n--
		//}

		if nums1[m] == nums2[n] {
			ans = append(ans, nums2[n])
			m--
			n--
		}else if nums1[m] > nums2[n] {
			m--
		}else {
			n--
		}
	}

	return ans
}

// Follow up
