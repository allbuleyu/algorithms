package prob1335

func minDifficulty(jobDifficulty []int, d int) int {
	return helpTopDown(jobDifficulty, d)
}

var maxVal = 1 << 31
var memory [][]int
var hardestJobRemaining []int

func helpTopDown(jobDifficulty []int, d int) int {
	if len(jobDifficulty) < d {
		return -1
	}

	memory = make([][]int, d+1)
	for i := 1; i <= d; i++ {
		memory[i] = make([]int, len(jobDifficulty))
	}

	hardestJobRemaining = make([]int, len(jobDifficulty))
	hardest := 0
	for i := len(jobDifficulty) - 1; i >= 0; i-- {
		hardest = max(hardest, jobDifficulty[i])
		hardestJobRemaining[i] = hardest
	}

	return topDown(jobDifficulty, 0, d)
}

func topDown(jobDifficulty []int, start, d int) int {
	if d == 0 {
		return hardestJobRemaining[start]
	}

	if memory[d][start] > 0 {
		return memory[d][start]
	}

	ans := maxVal
	todayDifficult := 0
	for i := start; i < len(jobDifficulty)-(d-1); i++ {
		todayDifficult = max(todayDifficult, jobDifficulty[i])
		ans = min(ans, topDown(jobDifficulty, i+1, d-1)+todayDifficult)
	}

	memory[d][start] = ans

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
