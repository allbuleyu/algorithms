package prob0078

func subsets(nums []int) [][]int {
	//res := make([][]int, 0)
	//cur := make([]int, 0)
	//
	//helper(&res, &cur, nums, 0)
	//return  res

	return bitManipulation(nums)
}

// recursion
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

// iterate
// 迭代方法还没懂
func helperIterate()  {

}

// 速度比递归稍慢
func bitManipulation(nums []int) [][]int {
	n := len(nums)
	resLen := 1 << uint(n)
	res := make([][]int, resLen)

	for i := range res {
		res[i] = make([]int, 0)
	}

	for i := 0; i < resLen; i++ {
		for j := 0; j < n; j++ {
			if i >> uint(j) & 1 == 1 {
				res[i] = append(res[i], nums[j])
			}

		}
	}

	return res
}