package prob1143

func longestCommonSubsequence(text1 string, text2 string) int {
	return helpTopDown(text1, text2)
}

func helpTopDown(t1, t2 string) int {
	memory = make([][]int, len(t1))
	for i := 0; i < len(t1); i++ {
		sub := make([]int, len(t2))
		for j := 0; j < len(t2); j++ {
			sub[j] = -1
		}
		memory[i] = sub
	}
	return topDown(t1, t2, len(t1)-1, len(t2)-1)
}

var memory [][]int

func topDown(t1, t2 string, x, y int) int {
	if x < 0 || y < 0 {
		return 0
	}

	ans := memory[x][y]
	if ans != -1 {
		return ans
	}

	if t1[x] != t2[y] {
		ans = max(topDown(t1, t2, x-1, y), topDown(t1, t2, x, y-1))
	} else {
		ans = topDown(t1, t2, x-1, y-1) + 1
	}

	memory[x][y] = ans
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
