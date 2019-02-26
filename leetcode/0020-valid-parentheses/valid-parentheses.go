package prob0020


func isValid(s string) bool {
	n := len(s)
	stack := make([]byte, n)
	ptr := 0

	for i := 0; i < n; i++ {
		v := s[i]
		switch v {
		case '(':
			stack[ptr] = v+1
			ptr++
		case '[', '{':
			stack[ptr] = v+2
			ptr++
		case ')', ']', '}':
			if ptr > 0 && stack[ptr-1] == v {
				ptr--
			}else {
				return false
			}

		}
	}


	return ptr == 0
}