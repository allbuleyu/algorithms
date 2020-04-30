package _530_week_03_3304

//Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
//(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
//
//You are given a target value to search. If found in the array return its index, otherwise return -1.
//
//You may assume no duplicate exists in the array.
//
//Your algorithm's runtime complexity must be in the order of O(log n).
//
//Example 1:
//
//Input: nums = [4,5,6,7,0,1,2], target = 0
//Output: 4
//Example 2:
//
//Input: nums = [4,5,6,7,0,1,2], target = 3
//Output: -1

// 相对于普通的binary search, 这道题相当于多了一个偏移量
func search(nums []int, target int) int {
	return binarySearchOptimize(nums, target)
}

// 参考他人优秀代码,优化方案
func binarySearchOptimize(nums []int, target int) int {
	// 偏移量
	startIndex := findRotatedIndexOptimize(nums)

	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l+r)/2
		realMid := (mid + startIndex) % len(nums)
		if nums[realMid] == target {
			return realMid
		}

		if nums[realMid] < target {
			l = mid + 1
		}else {
			r = mid - 1
		}
	}

	return -1
}

func findRotatedIndexOptimize(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		mid := (l + r) / 2

		if nums[mid] > nums[r] {
			l = mid + 1
		}else {
			r = mid
		}
	}

	return l
}

// 自己的方案
func search1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	index := findRotatedIndex(nums, 0, len(nums)-1)

	if index == -1 {
		return binarySearch(nums, target, 0, len(nums)-1)
	}

	l := binarySearch(nums, target, 0, index)
	if l != -1 {
		return l
	}

	r := binarySearch(nums, target, index, len(nums)-1)
	if r != -1 {
		return r
	}

	return -1
}

func binarySearch(nums []int, target int, start, end int) int {
	mid := (start + end) / 2
	for start != end {
		if nums[mid] < target {
			start = mid+1
		}

		if nums[mid] > target {
			end = mid
		}

		if nums[mid] == target {
			return mid
		}

		mid = (start + end) / 2
	}

	if nums[start] == target {
		return start
	}

	return -1
}


func findRotatedIndex(nums []int, start, end int) int {
	if start == end {
		return -1
	}

	mid := (start + end) / 2
	if mid > 0 && nums[mid] < nums[mid-1] {
		return mid
	}

	if mid < len(nums) - 1 && nums[mid] > nums[mid+1] {
		return mid + 1
	}

	lIndex := findRotatedIndex(nums, start, mid)
	rIndex := findRotatedIndex(nums, mid+1, end)

	if lIndex != -1 {
		return lIndex
	}

	if rIndex != -1 {
		return rIndex
	}

	return -1
}

