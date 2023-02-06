package prob0215

// https://leetcode.com/problems/kth-largest-element-in-an-array/description/
// Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.
// For example,
// Given [3,2,1,5,6,4] and k = 2, return 5.
//
// Note:
// You may assume k is always valid, 1 ≤ k ≤ array's length.
//
// Credits:
// Special thanks to @mithmatt for adding this problem and creating all test cases.
// Next challenge 324, 414
func findKthLargest(nums []int, k int) int {
	return divideAndConquer(nums, k)
}

func divideAndConquer(nums []int, k int) int {
	k = len(nums) - k
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := partition(nums, lo, hi)
		if mid == k {
			return nums[mid]
		} else if mid < k {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return nums[hi]
}

func partition(nums []int, lo, hi int) int {
	x := nums[hi]
	j := lo

	for ; j < hi; j++ {
		if nums[j] < x {
			nums[lo], nums[j] = nums[j], nums[lo]
			lo++
		}
	}

	nums[lo], nums[hi] = nums[hi], nums[lo]
	return lo
}
