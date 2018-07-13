package ch2

import "fmt"

func InsertSort(a []int) []int {
	for j := 1; j < len(a); j++ {

		//for i := 0; i < j; i++ {
		//	if a[j] < a[i] {
		//		a[j], a[i] = a[i], a[j]  // 冒泡排序
		//		fmt.Println(a)
		//	}
		//}
		key := a[j]
		i := j-1
		for ; i >= 0 && a[i] > key; i-- {

			a[i+1] = a[i]

		}
		a[i+1] = key
		fmt.Println(a)
	}

	return a
}
