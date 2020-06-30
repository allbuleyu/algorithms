package prob0037

import "fmt"

func solveSudoku(board [][]byte)  {
	helper(&board, 0, 0)
}

func helper(board *[][]byte, row, column int) bool {
	if row == 9 {
		return true
	}

	if column == 9 {
		return helper(board, row+1, 0)
	}

	if (*board)[row][column] != '.' {
		return helper(board, row, column+1)
	}
	var j byte = 1
	for ; j <= 9; j++ {
		if isValid(board, row, column, j+48) {
			(*board)[row][column] = j+48

			if helper(board, row, column+1) {
				return true
			}

			(*board)[row][column] = '.'
		}
	}

	return false
}

func isValid(board *[][]byte, row, column int, curNum byte) bool {

	// check for row
	for i := 0; i < 9; i++ {
		if i == column {
			continue
		}

		if (*board)[row][i] == curNum {
			return false
		}
	}

	// check for column
	for i := 0; i < 9; i++ {
		if i == row {
			continue
		}

		if (*board)[i][column] == curNum {
			return false
		}
	}

	// check for sub box
	r := row - (row%3)
	c := column - (column%3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if r+i == row && c+j == column {
				continue
			}

			if (*board)[r+i][c+j] == curNum {
				return false
			}
		}
	}

	return true
}

var rowKey = "%d in row %d"
var columnKey = "%d in column %d"
var blockKey = "%d in block %d - %d"
func isValidOptimize(num byte, row, column int, seen map[string]bool) bool {

	if num == '.' {
		return true
	}

	if _, ok := seen[fmt.Sprintf(rowKey, num, row)]; ok {
		return false
	}

	if _, ok := seen[fmt.Sprintf(columnKey, num, column)]; ok {
		return false
	}

	// 这一步实在是没懂...
	if _, ok := seen[fmt.Sprintf(blockKey, num, row/3, column/3)]; ok {
		return false
	}

	seen[fmt.Sprintf(rowKey, num, row)] = true
	seen[fmt.Sprintf(columnKey, num, column)] = true
	seen[fmt.Sprintf(blockKey, num, row/3, column/3)] = true

	return true
}