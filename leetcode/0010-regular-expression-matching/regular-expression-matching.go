package prob0010

func isMatch(s string, p string) bool {
	return recursionOfficial(s, p)
}

// todo
//
//	if (pattern.length() >= 2 && pattern.charAt(1) == '*'){
//	           return (isMatch(text, pattern.substring(2)) ||
//	                   (first_match && isMatch(text.substring(1), pattern)));
//	       } else {
//	           return first_match && isMatch(text.substring(1), pattern.substring(1));
//	       }
func recursionOfficial(s, p string) bool {
	if p == "" {
		return s == ""
	}

	var firstMatch bool
	firstMatch = !(s == "") && (p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return recursionOfficial(s, p[2:]) || (firstMatch && recursionOfficial(s[1:], p))
	}

	return firstMatch && recursionOfficial(s[1:], p[1:])
}

// todo
func DynamicProgramming(s, p string) bool {
	dp := make([][]bool, len(s)+1)
	//for i := 0; i < len(s); i++ {
	//	dp[i] = make([]bool, len(p)+1)
	//}
	//dp[0][0] = true
	//
	//for i := len(s); i >= 0; i-- {
	//	for j := len(p) - 1; j >= 0; j-- {
	//		firstMatch := i < len(s) && (s[i] == p[j] || p[j] == '.')
	//
	//	}
	//}

	return dp[0][0]
}
