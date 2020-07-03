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
	return letterCombinations1(digits)
}

func recursion(digits string) []string {
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


var dialMap = map[string][]byte{
	"2":[]byte{
		'a','b','c',
	},
	"3":[]byte{
		'd','e','f',
	},
	"4":[]byte{
		'g','h','i',
	},
	"5":[]byte{
		'j','k','l',
	},
	"6":[]byte{
		'm','o','n',
	},
	"7":[]byte{
		'p','q','s',
	},
	"8":[]byte{
		't','u','v',
	},
	"9":[]byte{
		'w','x','y','z',
	},
}

func letterCombinations1(digits string) []string {
	ans := make([]string, 0)
	cur := make([]byte, 0)
	backtrack(digits, &ans, &cur, 0, len(digits))
	return ans
}

func backtrack(digits string, ans *[]string, cur *[]byte, start, n int) {
	if start == n {
		if len(*cur) == 0 {
			return
		}
		//
		c := make([]byte, len(*cur))
		copy(c, *cur)
		*ans = append(*ans, string(c))
		return
	}

	t, ok := dialMap[string(digits[start])]
	if !ok {
		backtrack(digits, ans, cur, start+1, n)
		return
	}

	for i := 0; i < len(t); i++ {
		*cur = append(*cur, t[i])
		backtrack(digits, ans, cur, start+1, n)
		*cur = (*cur)[:len(*cur)-1]
	}
}