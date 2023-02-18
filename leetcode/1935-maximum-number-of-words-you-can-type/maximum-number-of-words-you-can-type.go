package prob1935

func canBeTypedWords(text string, brokenLetters string) int {
	hs := [26]bool{}
	for i := 0; i < len(brokenLetters); i++ {
		hs[brokenLetters[i]-'a'] = true
	}
	res := 0
	brkCnt := 0
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			//if text[i] == ' ' || i == len(text)-1 { wrong!!! take attention
			if brkCnt == 0 {
				res++
			}

			brkCnt = 0
		} else if hs[text[i]-'a'] {
			brkCnt++
		}
	}

	if brkCnt == 0 {
		res++
	}

	return res
}
