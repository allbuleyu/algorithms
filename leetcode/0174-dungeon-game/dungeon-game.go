package prob0174

func calculateMinimumHP(dungeon [][]int) int {
	return helpDp(dungeon)
}

func helpDp(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for c := n-1; c >= 0; c-- {
		if c == n-1 {
			dungeon[m-1][c] = getCurHp(dungeon[m-1][c], 1)
		}else {
			dungeon[m-1][c] = getCurHp(dungeon[m-1][c], dungeon[m-1][c+1])
		}
	}

	for r := m-2; r >= 0; r-- {
		for c := n - 1; c >= 0; c-- {
			if c == n-1 {
				dungeon[r][c] = getCurHp(dungeon[r][c], dungeon[r+1][c])
			}else {
				dungeon[r][c] = getCurHp(dungeon[r][c], min(dungeon[r+1][c], dungeon[r][c+1]))
			}
		}
	}

	return dungeon[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func getCurHp(cur, nextStep int) int {
	if cur < 0 {
		return -cur+nextStep
	}

	if cur >= nextStep {
		return 1
	}

	return nextStep-cur
}

