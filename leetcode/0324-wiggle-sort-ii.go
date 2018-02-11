package leetcode

import (
	"sort"
)

//https://leetcode.com/problems/wiggle-sort-ii/description/
//Given an unsorted array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....
//
//Example:
//(1) Given nums = [1, 5, 1, 1, 6, 4], one possible answer is [1, 4, 1, 5, 1, 6].
//(2) Given nums = [1, 3, 2, 2, 3, 1], one possible answer is [2, 3, 1, 3, 1, 2].
//[1,1,2,2,3,3], [1,2,1,3,2,3]
//[1,1,1,4,5,6], [1,4,1,5,1,6]
//例:[1, 5, 1, 1, 6, 4, 7]
//0,1,2,3,4,5,6, 0,4,1,5,2,6,3
//[1,1,1,4,5,6,7], [1,5,1,6,1,7,4]
//Note:
//You may assume all input has valid answer.
//
//Follow Up:
//Can you do it in O(n) time and/or in-place with O(1) extra space?
//
//Credits:
//Special thanks to @dietpepsi for adding this problem and creating all test cases.

func wiggleSort(nums []int)  {
	var splitLen1 int
	if len(nums) % 2 == 0 {
		splitLen1 = len(nums)/2
	}else {
		splitLen1 = (len(nums)+1)/2
	}

	sorted := sort.IsSorted(sort.IntSlice(nums))



}

func WiggleSort(nums []int)  {
	wiggleSort(nums)
}