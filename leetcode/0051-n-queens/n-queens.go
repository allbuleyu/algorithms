package prob0051

func solveNQueens(n int) [][]string {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	queens := make([][]string, 0)
	helpBacktracking(0, &board, &queens)

	return queens
}


func helpBacktracking(curRow int, board *[][]byte, queens *[][]string) {
	if curRow == len(*board) {
		appendQueens(queens, board)
		return
	}

	for i := 0; i < len(*board); i++ {
		if isValid(board, curRow, i) {
			(*board)[curRow][i] = 'Q'
			helpBacktracking(curRow+1, board, queens)
			(*board)[curRow][i] = '.'
		}
	}
}

func appendQueens(queens *[][]string, board *[][]byte) {
	n := len(*board)
	ap := make([]string, n)
	for i := 0; i < n; i++ {
		ap[i] = string((*board)[i])
	}

	*queens = append(*queens, ap)
}


func isValid(board *[][]byte, row, column int) bool {
	n := len(*board)
	// check for column
	for i := 0; i < row; i++ {
		if (*board)[i][column] == 'Q' {
			return false
		}
	}

	// check top left
	for i, j := row-1, column-1; i >= 0 && j >= 0; {
		if (*board)[i][j] == 'Q' {
			return false
		}

		i--
		j--
	}

	// check top right
	for i, j := row-1, column+1; i>=0 && j < n; {
		if (*board)[i][j] == 'Q' {
			return false
		}

		i--
		j++
	}

	return true
}