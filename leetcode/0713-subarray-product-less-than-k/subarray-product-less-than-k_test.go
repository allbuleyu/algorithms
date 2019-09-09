package prob0713

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_numSubarrayProductLessThanK(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		k int
		ans int
	}{
		//{[]int{10, 5, 2, 6},100,8},
		{[]int{10,1,2}, 10, 3},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, numSubarrayProductLessThanK(tc.input, tc.k), "输入:%v", tc)
	}
}

func Test_slidingWindowExt(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		k int
		ans [][]int
	}{
		//[10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
		{[]int{10, 5, 2, 6},100,[][]int{
			[]int{10}, []int{10, 5}, []int{5}, []int{5, 2}, []int{2}, []int{5, 2, 6}, []int{2, 6}, []int{6},
		}},
		{[]int{10,1,2}, 10, [][]int{
			[]int{1},[]int{1,2},[]int{2},
		}},
		{[]int{10,1}, 10, [][]int{
			[]int{1},
		}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, slidingWindowExt(tc.input, tc.k), "输入:%v", tc)
	}
}