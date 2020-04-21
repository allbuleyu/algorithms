package prob0074

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	if len(matrix) == 1 {
		return binarySearch(matrix[0], target, 0, len(matrix[0])-1)
	}

	index := 0
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] > target {
			index = i-1
			break
		}
		index = i
	}

	return binarySearch(matrix[index], target, 0, len(matrix[index]) -1)
}

func binarySearch(nums []int, target, start, end int) bool {
	if start == end {
		return false
	}
	res := false
	mid := (start + end) / 2
	if nums[mid] < target {
		res = binarySearch(nums, target, mid+1, end)
	}else if nums[mid] > target {
		res = binarySearch(nums, target, start, mid-1)
	}else {
		res = true
	}

	return res
}