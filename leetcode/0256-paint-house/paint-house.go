package prob0256

import (
	"fmt"
	"math"
)

func minCost(costs [][]int) int {
	return bottomUp(costs)
}

func bottomUp(costs [][]int) int {
	var r, b, g int

	for i := 0; i < len(costs); i++ {
		c := costs[i]
		r, b, g = min(b, g)+c[0], min(r, g)+c[1], min(r, b)+c[2]
	}

	return min(r, min(b, g))
}

var gCosts [][]int
var hs map[string]int

func helpTopDown(costs [][]int) int {
	gCosts = costs

	hs = make(map[string]int)
	return topDown(0, -1, 0)
}

func topDown(row, color, sum int) int {
	if row == len(gCosts) {
		return sum
	}

	hKey := fmt.Sprintf("%d-%d", row, color)
	if v, ok := hs[hKey]; ok {
		return v
	}

	ans := math.MaxInt32
	for i := 0; i < len(gCosts[0]); i++ {
		if i == color {
			continue
		}

		ans = min(ans, topDown(row+1, i, sum+gCosts[row][i]))
	}
	hs[hKey] = ans

	return ans
}

func bruteForce(row, column, sum int) int {
	if row == len(gCosts) {
		return sum
	}

	ans := math.MaxInt32
	for i := 0; i < len(gCosts[0]); i++ {
		if i == column {
			continue
		}

		ans = min(ans, bruteForce(row+1, i, sum+gCosts[row][i]))
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
