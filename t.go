package main

import (
	"encoding/base64"
	"fmt"
	"reflect"
)

func main() {

	b64 := base64.RawURLEncoding
	str := make([]byte, 1000)
	b64.Decode(str,[]byte("8IKW23gHgXTh0ri/tpfRyQ=="))

	fmt.Printf(string(str))
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



