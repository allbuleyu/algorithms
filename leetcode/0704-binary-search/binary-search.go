package prob0704

func search(nums []int, target int) int {
	return recursionHelp(nums, target, 0, len(nums)-1)
}

func recursionHelp(nums []int, target, left, right int) int {
	middle := (left + right)/2
	if nums[middle] == target {
		return middle
	}

	if left == right {
		return -1
	}

	if nums[middle] < target {
		return recursionHelp(nums, target, middle, right)
	}else {
		return recursionHelp(nums, target, left, middle)
	}
}