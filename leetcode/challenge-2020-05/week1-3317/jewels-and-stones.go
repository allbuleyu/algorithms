package week1_3317


func numJewelsInStones(J string, S string) int {
	m := make(map[int32]bool)
	res := 0
	for _, j := range J {
		m[j] = true
	}

	for _, s := range S {
		if _, ok := m[s];ok {
			res++
		}
	}

	return res
}
