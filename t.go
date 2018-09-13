package main

import (
	"github.com/allbuleyu/algorithms/introductionToAlgorighms/ch2"
	"fmt"
)

type myStruct struct {
	S string
	I int
	q string
}

//var arr []int = []int{2,4,5,6,1,2,3,6}
var arr []int = []int{2,8,7,1, 3, 5,6, 4}

func main() {

	s1 := "123"
	fmt.Println(&s1, s1)

	s1 = s1[0:1] + " " + s1[1:]
	fmt.Println(&s1, s1)


	i1 := 123
	i2 := 345

	fmt.Println(&i1, i1)

	i1 = i2
	fmt.Println(&i1, i1)


	return

	//ch2.MergeSort(arr, 0, len(arr))

	//ch2.MergeSortFor(arr)

	ch2.QuickSort(arr, 0, len(arr))
	fmt.Println(arr)





}




