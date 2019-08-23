package prob0026

func removeDuplicates(nums []int) int {
	return optimize(nums)
}

// two pointers solution
func optimize(nums []int) int {
	if len(nums) == 0 {
		return 0
	}


	var i, j = 0, 1

	for ; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i+1
}

func old(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var n, tmp int

	n=1
	tmp = nums[0]
	for _, v := range nums[1:] {
		if tmp != v {
			n++
			tmp = v
			nums[n-1] = v
		}
	}

	return n
}