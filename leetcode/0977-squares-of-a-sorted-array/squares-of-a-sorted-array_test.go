package prob0977

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortedSquares(t *testing.T) {

	tcs := []struct{
		input []int

		ans []int
	}{
		{
			[]int{-4,-1,0,3,10},[]int{0,1,9,16,100},
		},
		{
			[]int{-7,-3,2,3,11},[]int{4,9,9,49,121},
		},
	}

	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, sortedSquares(tc.input), "输入:%v", tc)
	}
}