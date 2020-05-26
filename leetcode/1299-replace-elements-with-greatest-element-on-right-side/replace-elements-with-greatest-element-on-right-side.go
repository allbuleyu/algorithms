package prob1299

import "sort"

//Given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.
//
//After doing so, return the array.
//
//
//
//Example 1:
//
//Input: arr = [17,18,5,4,6,1]
//Output: [18,6,6,6,1,-1]
//
//
//Constraints:
//
//1 <= arr.length <= 10^4
//1 <= arr[i] <= 10^5

func replaceElements(arr []int) []int {
	sArr := make([]int, len(arr))
	copy(sArr, arr)
	sort.Ints(sArr)

	start := 0
	for i, si := 0, len(arr)-1; i < len(arr)-1; i++ {
		if arr[i] == sArr[si] {
			for ii := i-1; ii >= 0 && ii >= start; ii-- {
				arr[ii] = sArr[si]
			}

			start = i
			si--
		}
	}

	arr[len(arr)-1] = -1

	return arr
}