package prob0240

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	return helpZigzag(matrix, target)
	//return divideAndConquer(&matrix, target, 0, len(matrix)-1, 0, len(matrix[0])-1)
}

// 因为矩阵是已排序的
// 所以把搜索的点置于左下,或者右上.按值大小去搜索,时间复杂度min(m, n)
func helpZigzag(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	row, col := m-1, 0
	for row >= 0 && col < n {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			col++
		} else {
			row--
		}
	}

	return false
}

// 可以把矩阵划分为4个子矩阵
// 两个矩阵可能包含target, 2个矩阵则绝对不包含target
// 如果target < mart[up][left] 或者 target > mart[down][right] 那么这两个矩阵就绝对不可能包含target
// 即 因为是有序的数组, 比最小的小,比最大的大.那么排除两个
// 怎么剔除两个不可能的矩阵,从上往下(从小往大)移动,直到到边界row==down,或者target > mart[row][mid]
func divideAndConquer(martrix *[][]int, target, up, down, left, right int) bool {
	if up > down || left > right {
		return false
	} else if target < (*martrix)[up][left] || target > (*martrix)[down][right] {
		return false
	}

	mid := left + (right-left)/2

	row := up
	for row <= down && (*martrix)[row][mid] <= target {
		if target == (*martrix)[row][mid] {
			return true
		}
		row++
	}

	return divideAndConquer(martrix, target, up, down, left, mid-1) || divideAndConquer(martrix, target, up, row-1, left+1, right)
}

const used = -1e9 - 1

func helpDacDfs(matrix [][]int, target int) bool {
	return dacDfs(matrix, target, 0, 0, len(matrix), len(matrix[0]))
}

func dacDfs(matrix [][]int, target, i, j, m, n int) bool {
	if i >= m || j >= n || matrix[i][j] == used {
		return false
	}

	if matrix[i][j] == target {
		return true
	}
	matrix[i][j] = used

	return dacDfs(matrix, target, i+1, j, m, n) || dacDfs(matrix, target, i, j+1, m, n)
}

func helpDacDfsBinarySearch(matrix [][]int, target int) bool {
	return dacDfsBinarySearch(matrix, target, 0, len(matrix)-1, 0, len(matrix[0])-1)
}

func dacDfsBinarySearch(matrix [][]int, target, up, down, left, right int) bool {
	if up < 0 || up > down || left < 0 || left > right {
		return false
	}

	hMid, vMid := (up+down)/2, (left+right)/2
	if matrix[hMid][vMid] == used {
		return false
	}

	if matrix[hMid][vMid] == target {
		return true
	}

	matrix[hMid][vMid] = used

	if matrix[hMid][vMid] < target {
		return dacDfsBinarySearch(matrix, target, hMid+1, down, left, right) || dacDfsBinarySearch(matrix, target, up, down, vMid+1, right)
	}

	return dacDfsBinarySearch(matrix, target, up, hMid, left, right) || dacDfsBinarySearch(matrix, target, up, down, left, vMid)
}
