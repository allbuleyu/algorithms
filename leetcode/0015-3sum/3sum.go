package prob0015

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	var i, lo, hi int

	for ; i < len(nums)-1; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		lo, hi = i+1, len(nums)-1
		s := 0 - nums[i]
		for lo < hi {
			if nums[lo] + nums[hi] == s {
				ans = append(ans, []int{nums[i], nums[lo], nums[hi]})

				// skip all duplicate item
				for lo < hi && nums[lo] == nums[lo+1] {
					lo++
				}

				// skip all duplicate item
				for lo < hi && nums[hi] == nums[hi-1] {
					hi--
				}

				lo++
				hi--
			}else if nums[lo] + nums[hi] < s {
				lo++
			}else {
				hi--
			}
		}
	}




	return ans
}

func resolution1(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	var left, mid, right int
	right = len(nums)-1

	for ; left < len(nums)-1; left++ {
		if nums[left] > 0 {
			break
		}

		if left > 0 && nums[left-1] == nums[left] {
			continue
		}

		mid = left+1
		for mid < right {
			s := nums[left] + nums[mid] + nums[right]

			if s < 0 {
				mid++
			}else if s > 0 {
				right--
			}else {
				ans = append(ans, []int{nums[left], nums[mid], nums[right]})

				mid++

				// skip duplicate items
				for mid < right && nums[mid] == nums[mid-1] {
					mid++
				}

				right--
				// skip duplicate items
				for mid < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}


	}

	return ans
}
