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

		ans += j-i+1
	}

	return ans
}