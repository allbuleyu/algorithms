package main

import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	head := kit.CreateNodes(nums)

	tnums := kit.TransferNodes(head)

	fmt.Println(tnums)

}




