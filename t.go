package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(2%4, 4%2, 10%3)
	return

	nums := []int{1, 2, 3, 4, 5}
	sliceNums := make([]int, len(nums))
	for i := range sliceNums {
		sliceNums[i] = nums[i]
	}

	cutSlice(nums)
	cutSlice(sliceNums)
	fmt.Printf("first cut, nums=%v, slice=%v \n", nums, sliceNums)

	t := reflect.TypeOf(&nums)
	st := reflect.TypeOf(&sliceNums)

	fmt.Printf("nums type is: %s \n", t.Elem().String())
	fmt.Printf("snums type is: %s \n", st.Elem().String())

	cutSliceByPoint(&nums)
	cutSliceByPoint(&sliceNums)
	fmt.Printf("second cut, nums=%v, slice=%v \n", nums, sliceNums)

}

func cutSlice(nums []int) {
	nums[1] = 100
	nums = nums[1:]
}

func cutSliceByPoint(nums *[]int) {
	*nums = (*nums)[1:]
}



