package prob0035

func searchInsert(nums []int, target int) int {
	res := 0
	for res < len(nums) {
		if nums[res] >= target {
			break
		}

		res++
	}

	return res
}