package prob0053

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return divideAndConquer(nums, 0, len(nums)-1)
}

func dynamicPrograming(nums []int) int {

	return 0
}

// 分治法
func divideAndConquer(nums []int, left,right int) int {
	if left == right {
		return nums[left]
	}

	var leftMaxSum, rightMaxSum, middleMaxSum int

	middle := (right+left)/2

	leftMaxSum = divideAndConquer(nums, left, middle)
	rightMaxSum = divideAndConquer(nums, middle+1, right)

	var leftSubMaxSum, rightSubMaxSum int = int(math.MinInt32), int(math.MinInt32)
	var leftSubSum, rightSubSum int
	for i := middle;i>=left;i-- {
		leftSubSum += nums[i]
		if leftSubMaxSum < leftSubSum {
			leftSubMaxSum = leftSubSum
		}
	}

	for i := middle+1; i <= right; i++ {
		rightSubSum += nums[i]
		if rightSubMaxSum < rightSubSum {
			rightSubMaxSum = rightSubSum
		}
	}

	middleMaxSum = leftSubMaxSum + rightSubMaxSum

	return max(leftMaxSum, rightMaxSum, middleMaxSum)
}

func max(a, b, c int) int {
	m := a
	if m < b {
		m = b
	}

	if m < c {
		m = c
	}

	return m
}