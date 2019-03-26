package prob0784

func letterCasePermutation(S string) []string {
	res := make([]string, 0, len(S))
	cur := make([]byte, len(S))

	for i := range S {
		cur[i] = S[i]
	}

	backtrack(&res, &cur, S, 0)

	return res
}

func backtrack(res *[]string, cur *[]byte, s string, start int) {
	if start == len(s) {
		*res = append(*res, string(*cur))
		return
	}

	backtrack(res, cur, s, start+1)

	if s[start] >= 'a' && s[start] <= 'z' {
		(*cur)[start] = s[start]+'A'-'a'
	}

	backtrack(res, cur, s, start+1)
}