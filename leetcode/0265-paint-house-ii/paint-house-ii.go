package prob0265

func minCostII(costs [][]int) int {
	return bottomUp(costs)
}

// n*k 的思路
// 因为题意是要找到最小的花费,两个相邻的房子不能相同颜色,所以只需要找到1st 小和2nd小的数值
// [
//
//		[1,2,3],
//	 [6,5,4]
//
// ]
func bottomUp(costs [][]int) int {
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
