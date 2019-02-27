package prob0746

func minCostClimbingStairs(cost []int) int {
<<<<<<< HEAD

}

func dynamicProgramming(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = cost[0]
	for i := 1; i<len(cost); i++ {

	}
=======
	return dynamicProgramming(cost)
}

func dynamicProgramming(cost []int) int {
	var f1, f2 int
	for i := 0; i < len(cost); i++ {
		f0 := cost[i] + min(f1, f2)
		f1, f2 = f0, f1
	}

	return min(f1, f2)
}

// todo
func recursion(cost []int) int {

	return 0
}


func min(a, b int) int {
	if a > b {
		return b
	}

	return a
>>>>>>> 34cfd67db642458eb59534443a215c006677fcca
}