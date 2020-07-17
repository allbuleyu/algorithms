package prob0054

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_spiralOrder(t *testing.T) {

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
			},[]int{1,2,3,6,9,8,7,4,5},
		},
		{
			[][]int{
				[]int{1, 2, 3, 4},
			},[]int{1,2,3,4},
		},
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{7, 8, 9},
			},[]int{1,2,3,9,8,7},
		},
		{
			[][]int{
				[]int{1},
				[]int{2},
				[]int{3},
			},[]int{1,2,3},
		},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, spiralOrder(tc.input), "输入:%v", tc)
	}
}