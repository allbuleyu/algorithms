package prob0784

func letterCasePermutation(S string) []string {
	res := make([]string, 0, len(S))
	cur := make([]byte, 0)

	backtrack(&res, &cur, S, 0)

	return res
}

func backtrack(res *[]string, cur *[]byte, s string, start int) {
	for i := start; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			*cur = append(*cur, s[i])
			backtrack(res, cur, s, i+1)
			*cur = (*cur)[:i]

			*cur = append(*cur, (s[i]+'A'-'a'))
			backtrack(res, cur, s, i+1)
			*cur = (*cur)[:i]
		}else {
			*cur = append(*cur, s[i])
		}

		if len(*cur) == len(s) {
			*res = append(*res, string(*cur))
			return
		}
	}

}