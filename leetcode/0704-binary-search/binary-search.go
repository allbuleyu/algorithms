package prob0704

func search(nums []int, target int) int {
	return iterateHelp(nums, target)
	//return recursionHelp(nums, target, 0, len(nums)-1)
}

func recursionHelp(nums []int, target, left, right int) int {
	middle := (left + right)/2
	if nums[middle] == target {
		return middle
	}

	if left == right {
		return -1
	}

	if nums[middle] > target {
		return recursionHelp(nums, target, left,middle)
	}else {
		return recursionHelp(nums, target, left,middle)
	}
}

func iterateHelp(nums []int, target int) int {
	n := len(nums)
	pivot := 0
	left := 0
	right := n-1
	for left <= right {
		pivot = (left + right) / 2

		if nums[pivot] == target {
			return pivot
		}

		if nums[pivot] < target {
			right = pivot - 1
		}else {
			left = pivot+1
		}
	}

	return -1
}