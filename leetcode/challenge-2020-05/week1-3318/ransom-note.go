package week1_3318

//Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.
//
//Each letter in the magazine string can only be used once in your ransom note.
//
//Note:
//You may assume that both strings contain only lowercase letters.
//
//canConstruct("a", "b") -> false
//canConstruct("aa", "ab") -> false
//canConstruct("aa", "aab") -> true

type letterStore [26]int
func canConstruct(ransomNote string, magazine string) bool {
	ls := letterStore{}
	for _, s := range magazine {
		ls[s-'a']++
	}

	for _, s := range ransomNote {
		ls[s-'a']--

		if ls[s-'a'] < 0 {
			return false
		}
	}

	return true
}