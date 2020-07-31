package prob0252

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	return bySort(intervals)
}

func bySort(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}

		return false
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}

	return true
}

func bruteForce(intervals [][]int) bool {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			//if intervals[j][0] > intervals[i][0] && intervals[j][0] < intervals[i][1] {
			//	return false
			//}
			//
			//if intervals[j][1] > intervals[i][0] && intervals[j][1] < intervals[i][1] {
			//	return false
			//}
			//
			//if intervals[i][0] > intervals[j][0] && intervals[i][0] < intervals[j][1] {
			//	return false
			//}
			//
			//if intervals[i][1] > intervals[j][0] && intervals[i][1] < intervals[j][1] {
			//	return false
			//}

			if overlap(intervals[i], intervals[j]) {
				return false
			}
		}
	}

	return true
}

func overlap1(i1,i2 []int) bool {
	return (i1[0] >= i2[0] && i1[0] < i2[1]) || (i2[0] >= i1[0] && i2[0] < i1[1])
}

// 很好用的技巧啊
func overlap(i1,i2 []int) bool {
	return min(i1[1], i2[1]) > max(i1[0], i2[0])
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}