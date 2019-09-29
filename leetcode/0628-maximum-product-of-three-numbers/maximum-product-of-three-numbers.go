package prob0628

func maximumProduct(nums []int) int {
	return onePass(nums)
}

// 一次迭代的思路
// 找出2个最小的数,找出3个最大的数
//
func onePass(nums []int) int {
	var min1, min2, max1, max2, max3 int
	max1, max2, max3 = -1001, -1001, -1001
	min1, min2 = 1001, 1001

	for _, v := range nums {
		switch {
		case v > max1:
			max1, max2, max3 = v, max1, max2
		case v > max2:
			max2, max3 = v, max2
		case v > max3:
			max3 = v

		}

		switch {
		case v < min1:
			min1, min2 = v, min1
		case v < min2:
			min2 = v
		}
	}

	return max(min1*min2*max1, max1*max2*max3)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}



