package prob0226

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
			[]int{1,2,3},		// 又又吃瘪
			[]int{1,3,2},
		},
		{
			[]int{4,2,7,1,3,6,9},		// 又又吃瘪
			[]int{4,7,2,9,6,3,1},
		},

	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, kit.ToIntsOriginal(invertTree(kit.NewTrees(tc.input).Root)), "输入:%v", tc)
	}
}