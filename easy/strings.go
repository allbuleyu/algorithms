package easy

import "fmt"

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

//https://leetcode.com/problems/sort-characters-by-frequency/description/
//Given a string, sort it in decreasing order based on the frequency of characters.
//
//Example 1:
//
//Input:
//"tree"
//
//Output:
//"eert"
//
//Explanation:
//'e' appears twice while 'r' and 't' both appear once.
//So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
//Example 2:
//
//Input:
//"cccaaa"
//
//Output:
//"cccaaa"
//
//Explanation:
//Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
//Note that "cacaca" is incorrect, as the same characters must be together.
//Example 3:
//
//Input:
//"Aabb"
//
//Output:
//"bbAa"
//
//Explanation:
//"bbaA" is also a valid answer, but "Aabb" is incorrect.
//Note that 'A' and 'a' are treated as two different characters.
func frequencySort(s string) string {
	var Charcnt [256]int
	for _, c := range s {
		fmt.Println(int(c), string(byte(int(c))))
		Charcnt[int(c)]++
	}

	fmt.Println(Charcnt)

	return ""
}

func FrequencySort(s string) string {
	return frequencySort(s)
}