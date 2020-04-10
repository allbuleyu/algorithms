package challenge

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3284/

func isHappy(n int) bool {
	isRepeat := map[int]bool{}
	for n != 1 {
		isRepeat[n] = true
		n = squares(n)

		if _, ok := isRepeat[n]; ok {
			return false
		}
	}

	return true
}

func squares(n int) int {
	res := 0
	for n != 0 {
		m := n%10
		res += m*m

		n /= 10
	}

	return res
}
