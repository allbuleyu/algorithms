package prob0881

import "sort"

func numRescueBoats(people []int, limit int) int {
	return twoPointers(people, limit)
}

// Greedy(tow pointers)
func twoPointers(a []int, limit int) int {
	sort.Ints(a)
	i, j := 0, len(a) - 1
	for i <= j {
		if a[i] + a[j] <= limit {
			i++
		}

		j--
	}

	return len(a) - 1 - j
}