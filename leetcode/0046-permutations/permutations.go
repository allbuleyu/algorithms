package prob0046


func permute(nums []int) [][]int {
	return recursion(nums)
}

func recursion(nums []int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)

	recursionHelper(&res, &cur, nums, len(nums), 0)

	return res
}

func recursionHelper(res *[][]int, cur *[]int, nums []int, n, start int) {

	if len(*cur) == n {
		c := make([]int, 0)
		copy(c, *cur)
		*res = append(*res, c)
		return
	}


	for i := start % n; i < n; i++ {
		*cur = append(*cur, nums[i])
		curLen := len(*cur)

		recursionHelper(res, cur, nums, n, start+1)
		*cur = (*cur)[:i]

		if curLen == n {
			return
		}
	}

}

