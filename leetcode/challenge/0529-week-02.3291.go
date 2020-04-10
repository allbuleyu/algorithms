package challenge

//https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3291/

func backspaceCompare(S string, T string) bool {
	ls, lt := len(S)-1, len(T)-1

	sCounter := 0
	tCounter := 0
	for ls >=0 || lt >= 0 {

		for ;ls >=0; ls-- {
			if S[ls] == '#' {
				sCounter++
				continue
			}

			if sCounter > 0 {
				sCounter--
			}else {
				break
			}
		}


		for ; lt >= 0; lt-- {
			if T[lt] == '#' {
				tCounter++
				continue
			}
			
			if tCounter > 0 {
				tCounter--
			}else {
				break
			}
		}
		
		if ls >= 0 && lt >= 0 && S[ls] != T[lt] {
			return false
		}
		
		if (ls >= 0) != (lt >= 0) {
			return false
		}
		
		ls--
		lt--
	}
	
	return true
}
