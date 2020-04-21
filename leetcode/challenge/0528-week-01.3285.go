package challenge

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
	return helpDivideConquer(nums, 0, len(nums)-1)
}

func helpFun1(nums []int) int {

	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		
	}

	return maxSum
}

func helpDivideConquer(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	mid := (start+end)/2

	leftSum := helpDivideConquer(nums, start, mid)
	rightSum := helpDivideConquer(nums, mid+1, end)

	return maxThree(leftSum, rightSum, leftSum+rightSum)
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func maxThree(i, j, k int) int {
	if i > j {
		if i > k {
			return i
		}
	}

	if j > k {
		return j
	}

	return k
}