package prob0054


func spiralOrder(matrix [][]int) []int {
	return iteration(matrix)
}

// var directions = map[string][]int{
//     "right":[]int{0, 1},
//     "down":[]int{1, 0},
//     "left":[]int{0, -1},
//     "up":[]int{-1, 0},
// }

type node struct{
	val []int
	size []int	// m,n,mm,mn
	next *node
}

var right = &node{
	val:[]int{0, 1},
	size:[]int{0, -1, 0, 0},
	next:down,
}

var down = &node{
	val:[]int{1, 0},
	size:[]int{-1, 0, 0, 0},
	next:left,
}

var left = &node{
	val:[]int{0, -1},
	size:[]int{0, 0, 0, 1},
	next:up,
}

var up = &node{
	val:[]int{-1, 0},
	size:[]int{0, 0, 1, 0},
	next:right,
}

func iteration(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	mm, mn := 0, 0

	direction := right
	r, c := 0, 0
	ans := make([]int, 0, m*n)
	for {
		x, y := direction.val[0], direction.val[1]
		for r < m && c < n && r >= mm && c >= mn {
			ans = append(ans, matrix[r][c])
			if r == mm || r == m-1 || c == mn || c == n-1 {
				break
			}
			r, c = r+x, c+y
		}

		m, n, mm, mn = m+direction.val[0], n+direction.val[1],mm+direction.val[2], mn+direction.val[3]
		direction = direction.next

		if mm == m-1 {
			break
		}
	}

	return ans
}