package prob0090

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0)
	cur := make([]int, 0)

	helper(&res, &cur, nums, 0)
	return  res
}

// recursion
func helper(res *[][]int, cur *[]int, nums []int, start int) {
	c := make([]int, len(*cur))
	copy(c, *cur)
	*res = append(*res, c)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		*cur = append(*cur, nums[i])
		helper(res, cur, nums, i+1)
		*cur = (*cur)[:len(*cur)-1]
	}
	return
}