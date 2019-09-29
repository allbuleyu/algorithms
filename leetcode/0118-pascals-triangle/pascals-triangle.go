package prob0118

func generate(numRows int) [][]int {
	res := make([][]int, 0)

	for i := 0; i < numRows; i++ {
		subArr := make([]int, i+1)
		subArr[0] = 1
		subArr[i] = 1
		for j := 1; j < i;j++ {
			subArr[j] = res[i-1][j] + res[i-1][j-1]
		}

		res = append(res, subArr)
	}

	return res
}