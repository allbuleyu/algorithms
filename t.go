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

	intSlice := []int{1,2,3,4,5,5,6,7,8,9,10,11}
	//
	fmt.Println(leetcode.IsPossible(intSlice))

	//words := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	//fmt.Println(leetcode.TopKFrequentStr(words, 4), "i" < "love")
}
