package main

import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
)

func main() {
	bys := []byte{'0','1', '2', '3', '4', '5', '6', '7', '8', '9'}
	fmt.Println(bys)
	return

	nums := []int{1, 2, 3, 4, 5, 6}

	head := kit.CreateNodes(nums)

	tnums := kit.TransferNodes(head)

	fmt.Println(tnums)

}




