package prob1155

func numRollsToTarget(n int, k int, target int) int {
	if n*k < target {
		return 0
	}

	if n < k {
		return numRollsToTarget(k, n, target)
	}

	return bottomUp(n, k, target)
}

func bottomUp(n, k, target int) int {
	return 0
}
