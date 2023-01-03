package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeSortTopDown(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct {
		input, ans []int
	}{
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		//{
		//	[]int{2, 7, 11, 15},
		//	[]int{2, 7, 11, 15},
		//},
		//{
		//	[]int{1, 2, 3, 4, 5, 6},
		//	[]int{1, 2, 3, 4, 5, 6},
		//},
		//{
		//	[]int{6, 5, 2, 3, 1, 4},
		//	[]int{1, 2, 3, 4, 5, 6},
		//},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		mergeSortTopDown(tc.input)
		ast.Equal(tc.ans, tc.input, "输入:%v", tc)
	}
}
