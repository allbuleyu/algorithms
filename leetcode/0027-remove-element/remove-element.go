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

	//println(resIndex)
	//return 0

	return resIndex
}