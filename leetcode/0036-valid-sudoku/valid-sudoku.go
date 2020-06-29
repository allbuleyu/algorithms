package prob0036

import "fmt"

func isValidSudoku(board [][]byte) bool {
	seen := make(map[string]bool)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !isValidOptimize(&board, i, j, seen) {
				return false
			}
		}
	}

	//for i := 0; i < 9; i++ {
	//	for j := 0; j < 9; j++ {
	//		if !isValid(&board, i, j) {
	//			return false
	//		}
	//
	//	}
	//}

	return true
}

func isValid(board *[][]byte, row, column int) bool {
	curNum := (*board)[row][column]
	if curNum == '.' {
		return true
	}
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

func isValidOptimize(board *[][]byte, row, column int, seen map[string]bool) bool {
	num := (*board)[row][column]
	if num == '.' {
		return true
	}

	if _, ok := seen[fmt.Sprintf("%d in row %d", num, row)]; ok {
		return false
	}

	if _, ok := seen[fmt.Sprintf("%d in column %d", num, column)]; ok {
		return false
	}

	// 这一步实在是没懂...
	if _, ok := seen[fmt.Sprintf("%d in block %d - %d", num, row/3, column/3)]; ok {
		return false
	}

	seen[fmt.Sprintf("%d in row %d", num, row)] = true
	seen[fmt.Sprintf("%d in column %d", num, column)] = true
	seen[fmt.Sprintf("%d in block %d - %d", num, row/3, column/3)] = true

	return true
}

// 得票最高答案
//Set seen = new HashSet();
//    for (int i=0; i<9; ++i) {
//        for (int j=0; j<9; ++j) {
//            char number = board[i][j];
//            if (number != '.')
//                if (!seen.add(number + " in row " + i) ||
//                    !seen.add(number + " in column " + j) ||
//                    !seen.add(number + " in block " + i/3 + "-" + j/3))
//                    return false;
//        }
//    }
//    return true;
func isValidSudokuOptimize(board [][]byte) bool {
	seen := make(map[string]bool)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}

			if _, ok := seen[fmt.Sprintf("%d in row %d", num, i)]; ok {
				return false
			}

			if _, ok := seen[fmt.Sprintf("%d in column %d", num, j)]; ok {
				return false
			}

			// 这一步实在是没懂...
			if _, ok := seen[fmt.Sprintf("%d in block %d - %d", num, i/3, j/3)]; ok {
				return false
			}

			seen[fmt.Sprintf("%d in row %d", num, i)] = true
			seen[fmt.Sprintf("%d in column %d", num, j)] = true
			seen[fmt.Sprintf("%d in block %d - %d", num, i/3, j/3)] = true
		}
	}

	return true
}