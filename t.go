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

	//intSlice := []int{4,5,5,6}
	intSlice := []int{2,1,3,4,5,4, 7}
	fmt.Println(leetcode.LargestRectangleArea(intSlice))
	return
	//
	fmt.Println(leetcode.IsPossible(intSlice))

	//words := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	//fmt.Println(leetcode.TopKFrequentStr(words, 4), "i" < "love")
}
