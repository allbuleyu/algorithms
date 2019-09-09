package prob0713

func numSubarrayProductLessThanK(nums []int, k int) int {
	return slidingWindow(nums, k)
}

func slidingWindow(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	var i, j, n, product int
	n = len(nums)
	ans := 0
	product = 1

	for ; j < n; j++ {
		product *= nums[j]
		for product >= k {
			product /= nums[i]
			i++
		}

		// 最核心的代码,这一步懂了,移动窗口就非常的简单
		// 思考: 如果要返回的是结果的数组,而不是个数,如何处理?
		ans += j-i+1
	}

	return ans
}

// 思考: 如果要返回的是结果的数组,而不是个数,如何处理?
func slidingWindowExt(nums []int, k int) [][]int {
	if k <= 1 {
		return [][]int{}
	}

	var i, j, n, product int
	n = len(nums)
	ans := make([][]int, 0)
	product = 1

	for ; j < n; j++ {
		product *= nums[j]
		for product >= k {
			product /= nums[i]
			i++
		}

		// i <= j 才是正确的解
		if i <= j {
			ans = append(ans, nums[i:j+1])
		}

		for t := i+1; t <= j; t++ {
			ans = append(ans, nums[t:j+1])
		}
	}

	return ans
}