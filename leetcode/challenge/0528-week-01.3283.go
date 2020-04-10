package challenge

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3283/


func singleNumber(nums []int) int {
	res := 0
	for i := range nums {
		res ^= nums[i]
	}

	return res
}
