package prob0007

import (
	"math"
)

func reverse(x int) int {
	var res int

	for x != 0 {
		y := x % 10

		z := res * 10 + y

		if z > math.MaxInt32 || z < math.MinInt32 {
			return 0
		}

		res = z

		x = x / 10
	}


	return res
}

