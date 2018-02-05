// Given an array of 2n integers, your task is to group these integers into n pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.

// Example 1:
// Input: [1,4,3,2]

// Output: 4
// Explanation: n is 2, and the maximum sum of pairs is 4
//
package easy

import (
	"fmt"
	"sort"
)


func arrayPairSum(nums []int) int {
	sum := 0

	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i += 2 {
		sum += nums[i]
	}

	return sum
}

func moveZeroes(nums []int) {
	zeroNum := 0
	for _, num := range nums {
		if num != 0 {
			nums[zeroNum] = num

			zeroNum++
		}
	}

	for ; zeroNum < len(nums); zeroNum++ {
		nums[zeroNum] = 0
	}
}

func romanToInt(s string) int {
	var sum rune
	r := []rune(s)
	fmt.Println(r)
	for _, val := range r {
		sum += val
	}

	return int(sum)
}
