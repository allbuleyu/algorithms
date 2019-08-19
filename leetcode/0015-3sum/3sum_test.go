package prob0015


import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_threeSum(t *testing.T) {

	tcs := []struct{
		input []int
		ans [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				[]int{-1, 0, 1},
				[]int{-1, -1, 2},
			},
		},

	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, threeSum(tc.input), "输入:%v", tc)
	}
}