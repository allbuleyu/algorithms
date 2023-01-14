package prob0740

import "sort"

func deleteAndEarn(nums []int) int {
	return dpBottomUp(nums)
}

// worst case time complexity O(x^2).
func f1(nums []int) int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]] += nums[i]
	}

	res := 0
	for i := range nums {
		x, y := nums[i]-1, nums[i]+1

		sum := 0
		for k, v := range m {
			if k != x && k != y {
				sum += v
			}
		}

		res = max(res, sum)
	}

	return res
}

func dpBottomUp(nums []int) int {
	sort.Ints(nums)

	m := make(map[int]int)
	for _, v := range nums {
		m[v] += v
	}

	var x, y int
	for i := 0; i < len(nums); i++ {
		for i > 0 && i < len(nums) && nums[i] == nums[i-1] {
			i++
		}

		if i == len(nums) {
			break
		}

		if _, ok := m[nums[i]-1]; !ok {
			x, y = x+m[nums[i]], x
		} else {
			x, y = max(x, y+m[nums[i]]), x
		}
	}

	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
