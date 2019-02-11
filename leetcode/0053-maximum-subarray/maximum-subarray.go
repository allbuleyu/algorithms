package prob0053

import (
	"math"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	//return divideAndConquer(nums, 0, len(nums)-1)
	return dynamicPrograming(nums)
}

func regular(nums []int) int {
	min := int(math.MinInt32)
	var ctn, res int = 0, min
	for i := 0; i < len(nums); i++ {
		ctn += nums[i]

		if ctn > res {
			res = ctn
		}

		if ctn < 0 {
			ctn = 0
		}
	}

	return res
}

func dynamicPrograming(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]

	twoMax := func(i, j int) int {
		if i > j {
			return i
		}

		return j
	}
	for i := 1; i < len(nums); i++ {
		dp[i] = twoMax(nums[i], dp[i-1]+nums[i])
		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
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