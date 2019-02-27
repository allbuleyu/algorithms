package prob0027


func removeElement(nums []int, val int) int {
	n := len(nums)

	var resIndex, j int
	for ; j < n; j++ {
		if nums[j] != val {
			if resIndex != j {
				nums[resIndex] = nums[j]
			}

			resIndex++
		}
	}

	// why not resIndex+1  thin of line 11 resIndex++ step
	return resIndex
}