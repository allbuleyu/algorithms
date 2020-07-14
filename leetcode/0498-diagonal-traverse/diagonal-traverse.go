package prob0498

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	//ans := make([]int, 0, m*n)
	//recursion(&matrix, &ans, 1, m, n, 0, 0)
	//return ans

	return iteration(matrix, m, n)
}

// 还有一种方法,从matrix[0][0] -> matrix[0][n], 都朝下迭代,把奇数迭代的结果翻转,这里不写了
// direction 1向上, 2向下
func recursion(matrix *[][]int, ans *[]int, direction, m, n, row, col int) {
	if row == m || col == n {
		return
	}

	*ans = append(*ans, (*matrix)[row][col])

	if direction == 1 {
		nRow, nCol := row-1, col+1
		if row == 0 {
			nRow = 0
			direction = 2
		}

		if col == n - 1 {
			nCol = col
			nRow = row+1
			direction = 2
		}

		recursion(matrix, ans, direction, m, n, nRow, nCol)
		return
	}

	nRow, nCol := row+1, col-1
	if col == 0 {
		nCol = 0
		direction = 1
	}

	if row == m -1 {
		nRow = row
		nCol = col+1
		direction = 1
	}

	recursion(matrix, ans, direction, m, n, nRow, nCol)
}

func iteration(matrix [][]int, m, n int) []int {
	row, col := 0, 0
	nRow, nCol := 0, 0
	ans := make([]int, 0, m*n)
	direction := 1
	for {
		ans = append(ans, matrix[row][col])

		if direction == 1 {
			nRow, nCol = row-1, col+1
			if row == 0 {
				nRow = row
				direction=2
			}

			if col == n - 1 {
				nRow = row+1
				nCol = col
				direction=2
			}

		}else {
			nRow, nCol = row+1, col-1
			if col == 0 {
				nCol = 0
				direction = 1
			}

			if row == m -1 {
				nRow = row
				nCol = col+1
				direction = 1
			}

		}

		row = nRow
		col = nCol
		if row == m || col == n {
			break
		}
	}

	return ans
}