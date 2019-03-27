package prob0784

import "strings"

func letterCasePermutation(S string) []string {
	res := make([]string, 0, len(S))
	cur := make([]byte, len(S))

	for i := range S {
		cur[i] = S[i]
	}

	S = strings.ToLower(S)

	backtrack(&res, &cur, S, 0)

	return res
}

// time 129 ms
// space 126mb  233
func backtrack(res *[]string, cur *[]byte, s string, start int) {
	if start == len(s) {
		*res = append(*res, string(*cur))
		return
	}


	if s[start] >= 'a' && s[start] <= 'z' {
		(*cur)[start] = s[start]
		backtrack(res, cur, s, start+1)
		(*cur)[start] = s[start]+'A'-'a'
		backtrack(res, cur, s, start+1)
	}else {
		backtrack(res, cur, s, start+1)
	}

}