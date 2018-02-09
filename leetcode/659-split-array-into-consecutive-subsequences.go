package leetcode

import "math"

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
//Next challenge hard(45,295,778)
func isPossible(nums []int) bool {
	firstEle := nums[0]
	lastEle := nums[len(nums) - 1]
	var size int

	if math.Abs(float64(firstEle)) > math.Abs(float64(lastEle)) {
		size = int(math.Abs(float64(firstEle))) + 1
	}else {
		size = int(math.Abs(float64(lastEle))) + 1
	}
	intNum := make([]int, size)

	var maxSplitSize int
	for i := range nums {
		v := nums[i]
		if nums[i] < 0 {
			v = -nums[i]
		}

		intNum[v]++

		if intNum[v] > maxSplitSize {
			maxSplitSize = intNum[v]
		}
	}

	splitSlice := make([][]int, maxSplitSize)

	max := 0
	for i, v := range intNum {
		if i == 0 {
			continue
		}

		if max < v {
			max = v
		}

		for j := 0; j < v; j++ {
			if v < max {
				splitSlice[max-v] = append(splitSlice[max-v], i)
				l := len(splitSlice[max-v])
				if l > 1 && (abs(splitSlice[max-v][l-1]) - abs(splitSlice[max-v][l-2])) != 1 {
					return false
				}
			}else {
				splitSlice[j] = append(splitSlice[j], i)
				l := len(splitSlice[j])
				if l > 1 && (abs(splitSlice[j][l-1]) - abs(splitSlice[j][l-2])) != 1 {
					return false
				}
			}
		}

	}

	for i := range splitSlice {
		if len(splitSlice[i]) < 3 {
			return false
		}
	}

	return true
}

func isPossible1(nums []int) bool {
	m := map[int]int{}
	group := map[int]int{}

	for _, v := range nums {
		m[v]++
	}

	for _, v := range nums {
		if m[v] == 0 {
			continue
		}
		m[v]--
		if group[v-1] != 0 {	// 判断长度大于3的子串的合法性, !=0 表示是连续的 = 0 表示第一个串或者新的子串
			group[v-1]--
			group[v]++
		} else {				//
			if m[v+1] != 0 && m[v+2] != 0 {			// 判断新的连续的子串的合法长度是否 = 3
				m[v+1]--
				m[v+2]--
				group[v+2]++
			} else {
				return false
			}
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0 // return correctly abs(-0)
	}
	return x
}

func IsPossible(nums []int) bool {
	return isPossible1(nums)
}