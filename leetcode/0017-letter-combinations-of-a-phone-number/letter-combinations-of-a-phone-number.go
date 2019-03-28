package prob0017

var m  =  map[byte]string{
	'2':"abc",
	'3':"def",
	'4':"ghi",
	'5':"jkl",
	'6':"mno",
	'7':"pqrs",
	'8':"tuv",
	'9':"wxyz",
}


func letterCombinations(digits string) []string {
	res := make([]string, 0)

	n := len(digits)

	if n == 0 || digits == "" || digits == "1" {
		return res
	}

	input := make([]string, 0)
	for i := 0; i < n; i++{
		s, ok := m[digits[i]]
		if ok {
			input = append(input, s)
		}
	}


	helper(&res, &[]byte{}, input, 0)

	return res
}

func helper(res *[]string, cur *[]byte, input []string, start int)  {
	if start == len(input) {
		c := make([]byte, start)
		copy(c, *cur)
		*res = append(*res, string(c))
		return
	}

	i:=start
	for j := 0; j < len(input[i]); j++ {
		*cur = append(*cur, input[i][j])
		helper(res, cur, input, i+1)
		*cur = (*cur)[:i]
	}

	return
}