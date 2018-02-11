package main

import (
	"github.com/allbuleyu/algorithms/leetcode"
	"fmt"
)

func main() {
	//intSlice := make([]int, 0, 15)
	//for i := 1; i <= 5; i++ {
	//	for j := 0; j< i; j++{
	//		intSlice = append(intSlice, i)
	//	}
	//}

	intSlice := []int{1,5,1,1,6,4}
	leetcode.WiggleSort(intSlice)
	return
	//
	fmt.Println(leetcode.IsPossible(intSlice))

	//words := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	//fmt.Println(leetcode.TopKFrequentStr(words, 4), "i" < "love")
}
