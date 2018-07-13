package ch1

import "strconv"

// 1.1.16
func ExR1(x int) string {
	if x <= 0 {
		return ""
	}

	return ExR1(x-3) + toString(x) + ExR1(x-2) + toString(x)
}

func toString(x int) string {
	return strconv.Itoa(x)
}

// 1.1.18
func Mystery(a, b int) int {
	if b == 0 {
		return 0
	}

	if b % 2 == 0 {
		return Mystery(a+a, b/2)
	}

	return Mystery(a+a, b/2) + a
}

// 1.1.19
//var fm map[int]int = make(map[int]int)
var fm []int = make([]int, 100)
func F(x int) int {
	//f, ok := fm[x]
	//if ok {
	//	return f
	//}

	if fm[x] != 0 {
		return fm[x]
	}

	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	v := F(x-1) + F(x-2)
	fm[x] = v

	return v
}


// 1.1.20
func Ln(x int64) int64 {
	if x <= 1 {
		return 1
	}

	return Ln(x-1) * x
}