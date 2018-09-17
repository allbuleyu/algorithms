package prob0387

//https://leetcode.com/problems/first-unique-character-in-a-string/description/
//Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
//
//Examples:
//
//s = "leetcode"
//return 0.
//
//s = "loveleetcode",
//return 2.
//Note: You may assume the string contain only lowercase letters.
//Next chanllenge 451
func firstUniqChar(s string) int {
	var charCnt [256]int

	for _, c := range s{
		charCnt[int(c)]++
	}

	for i, c := range s {
		if charCnt[int(c)] == 1 {
			return i
		}
	}

	return -1
}

func FirstUniqChar(s string) int {
	return firstUniqChar(s)
}