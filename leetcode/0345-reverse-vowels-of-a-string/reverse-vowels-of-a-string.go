package prob0345

func reverseVowels(s string) string {
	ans := []byte(s)
	for i, j := 0, len(s) - 1; i < j;  {
		for i < j && !isVowel(ans[i]) {
			i++
		}
		for i < j && !isVowel(ans[j]) {
			j--
		}

		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}

	return string(ans)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}