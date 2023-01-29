package prob0276

func totalWays(i, k int) int {
	return bottomUp(i, k)
}

func bottomUp(i, k int) int {
	if i == 1 {
		return k
	}
	if i == 2 {
		return k * k
	}

	prevOne, prevTwo := k*k, k
	for x := 2; x < i; x++ {
		prevOne, prevTwo = (k-1)*(prevOne+prevTwo), prevOne
	}

	return prevOne
}

func helperTopDown(i, k int) int {
	memory = make([]int, i)

	return topDown(i, k)
}

var memory []int

func topDown(i, k int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return k
	}
	if i == 2 {
		return k * k
	}
	if memory[i] > 0 {
		return memory[i]
	}

	memory[i] = (k - 1) * (topDown(i-1, k) + topDown(i-2, k))
	return memory[i]
}
