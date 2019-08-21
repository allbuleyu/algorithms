package prob0016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	return twoPointer(nums, target)
}

func twoPointer(nums []int, target int) int {
	sort.Ints(nums)
	var ans, i, lo, hi, minimum int
	minimum = int(math.MaxInt32)

	for ; i < len(nums)-1; i++ {
		lo = i+1
		hi = len(nums)-1

		for lo < hi {
			s := nums[i]+nums[lo]+nums[hi]
			if s-target < 0 {
				lo++
			}else if s-target > 0 {
				hi--
			}else {
				return s
			}

			if minimum > abs(s-target) {
				minimum = abs(s-target)
				ans = s
			}
		}
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func abs(x int) int {
	if x < 0 {
		return 0-x
	}

	return x
}