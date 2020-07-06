package prob0487

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	zeroIndex := -1
	left := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		//if nums[i] == 0 {
		//	ans = max(ans, i-zeroIndex+left)
		//	left  = i-zeroIndex-1
		//	zeroIndex = i
		//}
		//
		//if i == len(nums) - 1 {
		//	if zeroIndex == -1 {
		//		ans = len(nums)
		//	}else {
		//		ans = max(ans, i-zeroIndex+left+1)
		//	}
		//}


		// 优化版,简洁美观
		if nums[i] == 0 {
			left  = zeroIndex+1
			zeroIndex = i
		}

		ans = max(ans, i-left+1)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}