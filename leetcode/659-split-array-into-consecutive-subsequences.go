package leetcode

//https://leetcode.com/problems/split-array-into-consecutive-subsequences/description/
//You are given an integer array sorted in ascending order (may contain duplicates), you need to split them into several subsequences, where each subsequences consist of at least 3 consecutive integers. Return whether you can make such a split.
//
//Example 1:
//Input: [1,2,3,3,4,5]
//Output: True
//Explanation:
//You can split them into two consecutive subsequences :
//1, 2, 3
//3, 4, 5
//Example 2:
//Input: [1,2,3,3,4,4,5,5]
//Output: True
//Explanation:
//You can split them into two consecutive subsequences :
//1, 2, 3, 4, 5
//3, 4, 5
//Example 3:
//Input: [1,2,3,4,4,5]
//Output: False
//Note:
//The length of the input is in range of [1, 10000]

func isPossible(nums []int) bool {
	var size int = nums[len(nums) - 1] + 1
	var intNum [size]int
	var maxSplitSize int
	for i := range nums {
		intNum[nums[i]]++
		if intNum[nums[i]] > maxSplitSize {
			maxSplitSize = intNum[nums[i]]
		}
	}

	splitSlice := make([][]int, maxSplitSize)

	for i, v := range intNum {
		if i == 0 {
			continue
		}

		for j := 0; j < v; j++ {
			splitSlice[j] = append(splitSlice[j], i)
		}

	}

	for i := range splitSlice {
		if len(splitSlice[i]) < 3 {
			return false
		}
	}

	return true
}