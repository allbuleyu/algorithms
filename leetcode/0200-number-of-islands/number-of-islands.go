package prob0200

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	return bfs(grid)
}

var DIRECTIONS = [][]int{
	[]int{-1, 0},
	[]int{1, 0},
	[]int{0, -1},
	[]int{0, 1},
}

func bfs(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	queue := make([][]int, 0)
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				queue = append(queue, []int{i, j})
				grid[i][j] = '0'

				for len(queue) > 0 {
					point := queue[0]
					queue = queue[1:]
					row := point[0]
					col := point[1]

					for _, direction := range DIRECTIONS {
						r := direction[0] + row
						c := direction[1] + col
						if r < 0 || r == m || c < 0 || c == n || grid[r][c] == '0' {
							continue
						}

						queue = append(queue, []int{r, c})
						grid[r][c] = '0'
					}
				}
			}
		}
	}

	return ans
}