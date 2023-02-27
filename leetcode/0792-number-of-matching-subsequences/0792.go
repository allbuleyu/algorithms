package prob0792

func numMatchingSubseq(s string, words []string) (num int) {
	waiting := map[rune][]string{' ': words}
	for _, c := range " " + s {
		advance := waiting[c]
		delete(waiting, c)
		for _, word := range advance {
			if len(word) == 0 {
				num++
			} else {
				c := rune(word[0])
				waiting[c] = append(waiting[c], word[1:])
			}
		}
	}
	return
}
