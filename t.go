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

	intSlice := []int{1,1,1,2,2,3}

	fmt.Println(leetcode.TopKFrequent(intSlice, 2))
	return

	fmt.Println(leetcode.FrequencySort("leetcode"), 'z'+1)
}
