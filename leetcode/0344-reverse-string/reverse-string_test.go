package prob0344

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseString(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct {
		input []byte
		ans   string
	}{
		{[]byte("a"), "a"},
		{[]byte("ab"), "ba"},
		{[]byte("abc"), "cba"},
		{[]byte("abcde"), "edcba"},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		reverseString(tc.input)
		ast.Equal(tc.ans, string(tc.input), "输入:%v", tc)
	}
}

func Test_quickSort(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct {
		input []int
		ans   []int
	}{
		{[]int{1}, []int{1}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 3, 2, 1}, []int{1, 1, 1, 2, 3}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		quickSort(tc.input, 0, len(tc.input)-1)
		ast.Equal(tc.ans, tc.input, "输入:%v", tc)
	}
}

func quickSort(nums []int, i, j int) {
	if i >= j {
		return
	}

	p := partion(nums, i, j)
	quickSort(nums, i, p-1)
	quickSort(nums, p+1, j)
}

// 使用:将数组通过某一个值分为两组来理解.
func partion(nums []int, i, j int) int {
	x := nums[j]
	for k := i; k < j; k++ {
		if nums[k] <= x {
			swap(nums, i, k)
			i++
		}
	}

	// 为什么不用这个,因为这样i,k的距离就会非常远,
	// 很可能会用不到高速缓存.
	//k := j - 1
	//for i < k {
	//	if nums[i] <= x {
	//		i++
	//	} else {
	//		swap(nums, i, k)
	//		k--
	//	}
	//}

	swap(nums, i, j)
	return i
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
