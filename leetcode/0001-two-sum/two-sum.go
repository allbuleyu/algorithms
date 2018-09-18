package prob0001


//https://leetcode.com/problems/two-sum/description/
//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//Example:
//
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].


// 思路: a + b = c  => b = c - a
// 问题就转化为了找b的过程
func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if j, ok := m[complement]; ok {
			return []int{j, i}
		}

		m[nums[i]] = i
	}

	return []int{}
}

func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if j, ok := m[complement]; ok {
			return []int{i, j}
		}
	}

	return []int{}
}