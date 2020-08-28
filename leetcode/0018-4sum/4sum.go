package prob0018

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	n := len(nums)
	for i := 0; i < n-3; i++ {
		for i > 0 && i < n && nums[i] == nums[i-1] {
			i++
		}

		for j := i+1; j < n-2; j++ {
			for j > i+1 && j < n && nums[j] == nums[j-1] {
				j++
			}

			for lo, hi := j+1, n-1; lo < hi; {
				sum := nums[i] + nums[j] + nums[lo] + nums[hi]
				if sum < target {
					lo++
				}else if sum > target {
					hi--
				}else {
					ans = append(ans, []int{nums[i], nums[j], nums[lo], nums[hi]})
					lo++
					hi--

					for lo < hi && nums[lo] == nums[lo-1] {
						lo++
					}
				}
			}
		}
	}

	return ans
}


