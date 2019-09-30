package prob0566

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_matrixReshape(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input [][]int
		r,c int
		ans [][]int
	}{
		//{[]int{10, 5, 2, 6},100,8},
		{[][]int{[]int{1,2}, []int{3,4}}, 1, 4, [][]int{[]int{1,2,3,4}}},
		{[][]int{[]int{1,2}, []int{3,4}}, 2, 4, [][]int{[]int{1,2}, []int{3,4}}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, matrixReshape(tc.input, tc.r, tc.c), "输入:%v", tc)
	}
}