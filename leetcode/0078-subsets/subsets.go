package prob0078

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)

	helper(&res, &cur, nums, 0)
	return  res
}

func helper(res *[][]int, cur *[]int, nums []int, start int) {

	c := make([]int, len(*cur))
	copy(c, *cur)
	*res = append(*res, c)

	for i := start; i < len(nums); i++ {
		*cur = append(*cur, nums[i])
		helper(res, cur, nums, i+1)
		*cur = (*cur)[:len(*cur)-1]
	}
}