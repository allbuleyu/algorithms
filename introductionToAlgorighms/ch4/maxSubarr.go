package ch4

import "math"

func MaxSubArr01(a []int) (left, right, sum int) {
	if len(a) == 0 {
		return
	}

	sum = math.MinInt32
	subSum := a[0]

	for i := 0; i<len(a); i++ {
		if subSum+a[i] >= a[i] {
			subSum += a[i]
			right=i
		}else {
			subSum = a[i]

			left = i
			right= i
		}

		if sum < subSum {
			sum = subSum
		}
	}

	return
}
