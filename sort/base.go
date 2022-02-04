package sort

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func less(nums []int, i, j int) bool {
	return nums[i] < nums[j]
}

func exchange(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
