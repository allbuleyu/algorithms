package prob0844

func backspaceCompare(S string, T string) bool {
	return twoPointers(S, T)
}

func twoPointers(s, t string) bool {
	m, n := len(s) - 1, len(t) -1

	sCount := 0
	tCount := 0
	for m >= 0 || n >= 0 {
		for ;m >= 0;m-- {
			if s[m] == '#' {
				sCount++
				continue
			}

			if sCount == 0 {
				break
			}else {
				sCount--
			}

		}

		for ;n >= 0;n-- {
			if t[n] == '#' {
				tCount++
				continue
			}

			if tCount == 0 {
				break
			}else {
				tCount--
			}

		}

		if m >= 0 && n >= 0 && s[m] != t[n] {
			return false
		}

		if (m >= 0) != (n >= 0) {
			return false
		}

		m--
		n--
	}

	return true
}

func normal(s, t string) bool {
	bs := make([]byte, 0)
	bt := make([]byte, 0)

	bCount := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			bCount++
			continue
		}

		if bCount == 0 {
			bs = append(bs, s[i])
		}else {
			bCount--
		}
	}

	bCount = 0
	for i := len(t) - 1; i >= 0; i-- {
		if t[i] == '#' {
			bCount++
			continue
		}

		if bCount == 0 {
			bt = append(bt, t[i])
		}else {
			bCount--
		}
	}

	if string(bs) == string(bt) {
		return true
	}

	return false
}