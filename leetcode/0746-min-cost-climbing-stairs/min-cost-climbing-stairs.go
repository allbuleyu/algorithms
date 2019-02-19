package prob0746

func minCostClimbingStairs(cost []int) int {
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
}