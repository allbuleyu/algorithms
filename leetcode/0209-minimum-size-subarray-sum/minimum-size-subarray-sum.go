package prob0209

//https://leetcode.com/problems/minimum-size-subarray-sum/description/
//Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.
//
//For example, given the array [2,3,1,2,4,3] and s = 7,
//the subarray [4,3] has the minimal length under the problem constraint.
//
//click to show more practice.
//More practice:
//If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n).
//
//Credits:
//Special thanks to @Freezen for adding this problem and creating all test cases.
//next challenge 718
// 移动窗口
func minSubArrayLen(s int, nums []int) int {
	var subLen, start, sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		for sum >= s && start <= i {
			if subLen == 0 || i+1 - start < subLen {
				subLen = i+1 - start
			}
			sum -= nums[start]
			start++
		}
	}

	return subLen
}

func MinSubArrayLen(s int, nums []int) int {
	return minSubArrayLen(s, nums)
}