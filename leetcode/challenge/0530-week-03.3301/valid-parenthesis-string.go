package _530_week_03_3301

//Given a string containing only three types of characters: '(', ')' and '*', write a function to check whether this string is valid. We define the validity of a string by these rules:
//
//Any left parenthesis '(' must have a corresponding right parenthesis ')'.
//Any right parenthesis ')' must have a corresponding left parenthesis '('.
//Left parenthesis '(' must go before the corresponding right parenthesis ')'.
//'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
//An empty string is also valid.
//Example 1:
//Input: "()"
//Output: True
//Example 2:
//Input: "(*)"
//Output: True
//Example 3:
//Input: "(*))"
//Output: True
//Note:
//The string size will be in the range [1, 100].

func checkValidString(s string) bool {
	l := 0
	a := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
		}

		if s[i] == ')' {
			l--
		}
		
		if l < 0 {
			return false
		}
		
		if s[i] == '*' {
			a++
		}
	}


	return a >= l
}