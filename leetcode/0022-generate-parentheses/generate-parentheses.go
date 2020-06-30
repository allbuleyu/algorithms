package prob0022

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	cur := make([]byte, 0)
	backtrack(n, &ans, &cur, 0, 0)

	return ans
}

func backtrack(n int, ans *[]string, cur *[]byte, open, closed int) {
	if 2*n == len(*cur) {
		*ans = append(*ans, string(*cur))
	}

	if open < n {
		*cur = append(*cur, '(')
		backtrack(n, ans, cur, open+1, closed)

		*cur = (*cur)[:len(*cur)-1]
	}

	if closed < open {
		*cur = append(*cur, ')')
		backtrack(n, ans, cur, open, closed+1)
		*cur = (*cur)[:len(*cur)-1]
	}
}