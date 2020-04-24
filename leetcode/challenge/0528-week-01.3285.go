package challenge

import "math"

// https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/528/week-1/3285/


//Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
//
//Example:
//
//Input: [-2,1,-3,4,-1,2,1,-5,4],
//Output: 6
//Explanation: [4,-1,2,1] has the largest sum = 6.
//Follow up:
//
//If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
func maxSubArray(nums []int) int {
	return helpDp(nums)
	//return helpDivideConquer(nums, 0, len(nums)-1)
}

func helpDp(nums []int) int {
	var m, maxSum int
	m = nums[0]
	maxSum = nums[0]
	for i := 1; i < len(nums); i++ {
		m = max(m+nums[i], nums[i])
		maxSum = max(maxSum, m)
	}

	return maxSum
}

// 非常精彩的解法
func helpDivideConquer(nums []int, start, end int) int {
	if start > end {
		return math.MinInt32
	}

	m := start + (end-start)/2
	ml := 0
	mr := 0

	lmax := helpDivideConquer(nums, start, m-1)
	rmax := helpDivideConquer(nums, m+1, end)

	for i, sum := m-1, 0; i >= start; i-- {
		sum += nums[i]
		ml = max(sum, ml)
	}

	for i, sum := m+1, 0; i <= end; i++ {
		sum += nums[i]
		mr = max(sum, mr)
	}

	return max(max(lmax, rmax), ml+mr+nums[m])
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}