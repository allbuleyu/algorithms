package prob1122

func relativeSortArray(arr1 []int, arr2 []int) []int {
	return countingSort(arr1, arr2)
}

func countingSort(arr1 []int, arr2 []int) []int {
	var count [1001]int
	for _, v := range arr1 {
		count[v]++
	}

	var i int
	for _, v := range arr2 {
		for count[v] > 0 {
			arr1[i] = v
			i++

			count[v]--
		}
	}

	for j := 0; j < len(count); j++{
		for count[j] > 0 {
			arr1[i] = j
			i++
			count[j]--
		}
	}

	return arr1
}
