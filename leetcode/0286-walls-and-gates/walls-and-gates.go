package prob0286

import "math"

var DIRECTIONS  = [][]int{
	[]int{-1, 0},
	[]int{1, 0},
	[]int{0, -1},
	[]int{0, 1},
}

var EMPTY = math.MaxInt32
var WALL = -1
var GATE = 0

func wallsAndGates(rooms [][]int)  {
	if len(rooms) == 0 || len(rooms[0]) == 0 {
		return
	}
	bfs(rooms)
}

// 现阶段的queue的实现 queue[0], queue[:1] 速度真的感人,后续自己实现高效率的queue
func bfs(rooms [][]int) {
	m := len(rooms)
	n := len(rooms[0])
	queue := make([][]int, 0)

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if rooms[r][c] == GATE {
				queue = append(queue, []int{r, c})
			}
		}
	}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		row := point[0]
		col := point[1]
		for _, direction := range DIRECTIONS {
			r := row + direction[0]
			c := col + direction[1]
			if r < 0 || c < 0 || r >= m || c >= n || rooms[r][c] != EMPTY {
				continue
			}

			rooms[r][c] = rooms[row][col]+1
			queue = append(queue, []int{r, c})
		}
	}
}
