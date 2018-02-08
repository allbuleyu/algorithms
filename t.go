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

	//intSlice := []int{3,2,1,5,6,4}
	//
	//fmt.Println(leetcode.FindKthLargest(intSlice, 2), "i" < "love"[:1], "love"[:1])

	words := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	fmt.Println(leetcode.TopKFrequentStr(words, 4), "i" < "love")
}
