package prob0918

// todo:
func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	res := nums[0]
	sum := nums[0]
	beginIndex := 0

	for i := 1; i < 2*n; i++ {
		index := i % n
		if index == beginIndex {
			break
		}

		if nums[index] > sum+nums[index] {
			beginIndex = index

			sum = nums[index]
		}

		res = max(res, sum)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
