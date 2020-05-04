package week1_3316


func firstBadVersion(n int) int {
	l := 1
	r := n

	fV := -1
	for l <= r {
		mid := (l+r) / 2
		if !isBadVersion(mid) {
			l = mid + 1
		}else {
			r = mid - 1
			fV = mid
		}
	}

	return fV
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func isBadVersion(v int) bool {
	return true
}
