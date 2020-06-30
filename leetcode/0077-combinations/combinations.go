package prob0077


func combine(n int, k int) [][]int {

	ans := make([][]int, 0)
	cur := make([]int, 0)
	backtrack(n,k, 1, &ans, &cur)
	return ans
}

func backtrack(n, k, start int, ans *[][]int, cur *[]int) {
	if k == 0 {
		c := make([]int, len(*cur))
		copy(c, *cur)
		*ans = append(*ans, c)

		return
	}

	for i := start; i <= n+1-k; i++ {
		*cur = append(*cur, i)
		backtrack(n, k-1, i+1, ans, cur)
		*cur = (*cur)[:len(*cur)-1]
	}
}