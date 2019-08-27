package prob0977

func sortedSquares(A []int) []int {
	var i, j int
	j = len(A)-1

	ans := make([]int, j+1)
	for k := j;i <= j;k-- {
		if A[i] < 0 && square(A[i]) > square(A[j]) {
			ans[k] = square(A[i])
			i++
		}else {
			ans[k] = square(A[j])
			j--
		}

	}

	return ans
}

func abs(a int) int {
	if a < 0 {
		return 0-a
	}

	return a
}

func square(a int) int {
	return a * a
}