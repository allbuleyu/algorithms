package prob0498

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findDiagonalOrder(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct{
		input [][]int
		ans []int
	}{
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},[]int{1,2,4,7,5,3,6,8,9},
		},
		{
			[][]int{
				[]int{1, 2, 3, 4},
				[]int{4, 5, 6, 7},
				[]int{7, 8, 9, 10},
			},[]int{1,2,4,7,5,3,4,6,8,9,7,10},
		},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, findDiagonalOrder(tc.input), "输入:%v", tc)
	}
}