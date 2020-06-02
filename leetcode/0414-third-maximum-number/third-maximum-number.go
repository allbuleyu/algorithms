package prob0414

import (
	"math"
	"sort"
)

//https://leetcode.com/problems/third-maximum-number/description/
//Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).
//
//Example 1:
//Input: [3, 2, 1]
//
//Output: 1
//
//Explanation: The third maximum is 1.
//Example 2:
//Input: [1, 2]
//
//Output: 2
//
//Explanation: The third maximum does not exist, so the maximum (2) is returned instead.
//Example 3:
//Input: [2, 2, 3, 1]
//
//Output: 1
//
//Explanation: Note that the third maximum here means the third maximum distinct number.
//Both numbers with value 2 are both considered as second maximum.
//
//difficult: easy
func thirdMax(nums []int) int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}

	if len(m) < 3 {
		sort.Ints(nums)
		return nums[len(nums)-1]
	}

	r := make([]int, 0)
	for i := range m {
		r = append(r, i)
	}

	sort.Ints(r)
	return r[len(r) -  3]
}

func thirdMax1(nums []int) int {
	const kMax = 3
	hasValue := [kMax]bool{}
	maxValue := [kMax]int{}
	for _, v := range nums {
		for i := 0; i < kMax; i++ {
			if hasValue[i] {
				if v == maxValue[i] {
					break
				}
				if v > maxValue[i] {
					v, maxValue[i] = maxValue[i], v
				}
			} else {
				maxValue[i] = v
				hasValue[i] = true
				break
			}
		}
	}
	if hasValue[kMax-1] {
		return maxValue[kMax-1]
	}
	return maxValue[0]
}

func ThirdMax(nums []int) int {
	return thirdMax(nums)
}

func ffff(nums []int) int {
	f := math.MinInt64
	s := math.MinInt64
	t := math.MinInt64
	n := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] > f {
			f,s,t = nums[i], f, s
			n++
		}else if nums[i] > s && nums[i] < f {
			s, t = nums[i], s
			n++
		}else if nums[i] > t && nums[i] < s {
			t = nums[i]
			n++
		}
	}

	if n < 3 {
		return f
	}

	return t
}