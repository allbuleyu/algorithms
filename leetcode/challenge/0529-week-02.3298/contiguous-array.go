package _529_week_02_3298

//https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3298/
//Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.
//
//Example 1:
//Input: [0,1]
//Output: 2
//Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.
//Example 2:
//Input: [0,1,0]
//Output: 2
//Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
//Note: The length of the given binary array will not exceed 50,000.

func findMaxLength(nums []int) int {
	return helpExtraSpace(nums)
}

func helpHash(nums []int) int {
	m := make(map[int]int)

	// 从开头就有符合题意的数组的话,需要把m[0]=-1
	// m[1] = 0 || m[-1] = 0的情况
	m[0] = -1

	maxLen := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count--
		}else {
			count++
		}

		if pi, ok := m[count]; !ok {
			m[count] = i
		}else {
			maxLen = max(maxLen, i-pi)
		}
	}

	return maxLen
}

// 细细品
func helpExtraSpace(nums []int) int {
	extra := make([]int, len(nums)*2+1)

	for i := 0; i < len(extra); i++ {
		extra[i] = -2
	}
	extra[len(nums)] = -1

	maxLen := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count--
		}else {
			count++
		}

		if extra[count + len(nums)] >= -1 {
			maxLen = max(maxLen, i - extra[count + len(nums)])
		}else {
			extra[count + len(nums)] = i
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

