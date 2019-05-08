package prob0199

import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rightSideView(t *testing.T) {

	tcs := []struct{
		input []int
		ans []int
	}{
		{
			[]int{1,kit.Null,2,kit.Null,5,4,6,3},		// 又又吃瘪
			[]int{1,2,5,6,3},
		},
		{
			[]int{1,2,3,4},		// 吃瘪
			[]int{1,3,4},
		},
		{
			[]int{1,2,3,5,6},		// 又一次吃瘪
			[]int{1,3,6},
		},

		{
			[]int{1,2},
			[]int{1,2},
		},
		{
			[]int{1,2,3},
			[]int{1,3},
		},
		{
			[]int{1,2,3,kit.Null,5,kit.Null,4},
			[]int{1,3,4},
		},
	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, rightSideView(kit.NewTrees(tc.input).Root), "输入:%v", tc)
	}
}