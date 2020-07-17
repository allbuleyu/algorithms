package prob0054


func spiralOrder(matrix [][]int) []int {
	return iteration(matrix)
}

type node struct{
	val []int
	size []int	// m,n,mm,mn
	next *node
}

var right = &node{
	val:[]int{0, 1},
	size:[]int{0, 0, 1, 0},
	next:down,
}

var down = &node{
	val:[]int{1, 0},
	size:[]int{0, -1, 0, 0},
	next:left,
}

var left = &node{
	val:[]int{0, -1},
	size:[]int{-1, 0, 0, 0},
	next:up,
}

var up = &node{
	val:[]int{-1, 0},
	size:[]int{0, 0, 0, 1},
	next:nil,
}

func iteration(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	r1, r2, c1, c2 := 0, m-1, 0, n-1
	ans := make([]int, 0, m*n)
	for r1 <= r2 && c1 <= c2 {
		// direction right
		for c := c1; c <= c2; c++ {
			ans = append(ans, matrix[r1][c])
		}

		// direction down
		for r := r1+1; r <= r2; r++ {
			ans = append(ans, matrix[r][c2])
		}

		// 这里把单行,单列的排除
		if r1 < r2 && c1 < c2 {
			// direction left
			for c := c2-1; c >= c1; c-- {
				ans = append(ans, matrix[r2][c])
			}

			// direction up
			for r := r2-1; r >= r1+1; r-- {
				ans = append(ans, matrix[r][c1])
			}
		}

		r1++
		c1++
		r2--
		c2--
	}

	return ans
}

//func iteration(matrix [][]int) []int {
//	if len(matrix) == 0 || len(matrix[0]) == 0 {
//		return []int{}
//	}
//	m, n := len(matrix), len(matrix[0])
//	mm, mn := 0, 0
//
//	direction := right
//	up.next=right
//	r, c := 0, 0
//	ans := make([]int, 0, m*n)
//	for {
//		x, y := direction.val[0], direction.val[1]
//		for r < m && c < n && r >= mm && c >= mn {
//			ans = append(ans, matrix[r][c])
//			r, c = r+x, c+y
//		}
//
//
//		m, n, mm, mn = m+direction.size[0], n+direction.size[1],mm+direction.size[2], mn+direction.size[3]
//		direction = direction.next
//
//		if mm == m {
//			break
//		}
//	}
//
//	return ans
//}