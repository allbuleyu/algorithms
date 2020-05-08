package week1_3320


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

type letterStore [26]int
func firstUniqChar(s string) int {
	ls := letterStore{}
	for i := 0; i < len(s); i++ {
		ls[s[i] - 'a']++
	}

	for i := 0; i < len(s); i++ {
		if ls[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
