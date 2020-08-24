package prob0556

import (
	"math"
	"strconv"
)

func nextGreaterElement(n int) int {
	b := []byte(strconv.Itoa(n))

	i := len(b)-1
	firstLessIndex := i - 1
	for firstLessIndex >= 0 && b[firstLessIndex] >= b[i] {
		i = firstLessIndex
		firstLessIndex--
	}

	if firstLessIndex == -1 {
		return -1
	}

	nextGreaterIndex := firstLessIndex + 1
	for nextGreaterIndex < len(b) && b[nextGreaterIndex] > b[firstLessIndex] {
		nextGreaterIndex++
	}

	b[firstLessIndex], b[nextGreaterIndex-1] = b[nextGreaterIndex-1], b[firstLessIndex]

	reverse(b, firstLessIndex+1, len(b)-1)

	ans, _ := strconv.Atoi(string(b))
	if ans == n || ans > math.MaxInt32 {
		return -1
	}

	return ans
}

func reverse(b []byte, i, j int) {
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
}