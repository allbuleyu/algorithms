package prob0566

func matrixReshape(nums [][]int, r int, c int) [][]int {
	return normal(nums, r, c)
}

func normal(nums [][]int, r int, c int) [][]int {

	count := r * c
	// 矩阵的特性
	if len(nums) == 0 || count != len(nums) * len(nums[0]) {
		return nums
	}

	//res := make([][]int, 0)
	//subRes := make([]int, 0)
	//subLen := c
	//for _, subArr := range nums {
	//	for _, v := range subArr {
	//		subRes = append(subRes, v)
	//		subLen--
	//
	//		if subLen == 0 {
	//			res = append(res, subRes)
	//			subRes = make([]int, 0)
	//			subLen = c
	//		}
	//	}
	//}

	res :=  make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}

	count = 0
	for i := range nums {
		for j := 0; j < len(nums[i]); j++ {

			// 不用每次去判断数组的长度了
			res[count / c][count % c] = nums[i][j]
			count++
		}
	}

	return res
}