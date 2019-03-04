package main

import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
	"reflect"
)

func main() {
	var tmp byte
	tmp = '0'
	fmt.Println(reflect.TypeOf(&tmp).Elem().Name())
	return

	nums := []int{1, 2, 3, 4, 5, 6}

	head := kit.CreateNodes(nums)

	tnums := kit.TransferNodes(head)

	fmt.Println(tnums)

}




