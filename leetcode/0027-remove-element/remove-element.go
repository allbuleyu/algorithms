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

func helpFunc1(nums []int, val int) int {
	delNum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			delNum++
			continue
		}

		if delNum > 0 {
			nums[i-delNum] = nums[i]
		}
	}

	return len(nums) - delNum
}