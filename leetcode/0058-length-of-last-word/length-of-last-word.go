package prob0058

import "strings"

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	arr := strings.Split(s, " ")
	
	if len(arr)== 1 {
		return 1
	}
	
	return len(arr[len(arr)-1])
}