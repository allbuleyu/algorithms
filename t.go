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

	//ch2.MergeSort(arr, 0, len(arr))

	//ch2.MergeSortFor(arr)

	ch2.QuickSortDesc(arr, 0, len(arr))
	fmt.Println(arr)

}




