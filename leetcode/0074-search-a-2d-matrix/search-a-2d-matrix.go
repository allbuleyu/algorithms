package prob0074

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	return binSrch(matrix, target)
}

func helpZigzag(matrix [][]int, target int) bool {
	i := len(matrix)-1
	j := 0

	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] == target {
			return true
		}else if matrix[i][j] < target {
			j++
		}else {
			i--
		}
	}

	return false
}


func binSrch(matrix [][]int, target int) bool {
	left, right := 0, len(matrix[0])-1
	row := 0
	for row < len(matrix) {
		if matrix[row][left] <= target && matrix[row][right] >= target {
			break
		}
		row++
	}

	if row == len(matrix) {
		return false
	}

	for left <= right {
		mid := left + (right-left)/2
		if matrix[row][mid] == target {
			return true
		}else if matrix[row][mid] < target {
			left = mid + 1
		}else {
			right = mid - 1
		}
	}

	return false
}