package prob0011

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxArea(t *testing.T) {

	tcs := []struct{
		input []int
		ans int
	}{
		{
			[]int{1,8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			[]int{1},
			0,
		},
		{
			[]int{1,2},
			1,
		},
	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, maxArea(tc.input), "输入:%v", tc)
	}
}