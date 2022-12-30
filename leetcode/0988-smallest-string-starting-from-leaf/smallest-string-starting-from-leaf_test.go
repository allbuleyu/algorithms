package prob0988


import (
	"fmt"
	"algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_smallestFromLeaf(t *testing.T) {

	tcs := []struct{
		input []int
		ans string
	}{
		{
			[]int{0,1,2,3,4,3,4},		// 又又吃瘪
			"dba",
		},
		{
			[]int{25,1,3,1,3,0,2},		// 又又吃瘪
			"adz",
		},
		{
			[]int{2,2,1,kit.Null,1,0,kit.Null,0},		// 又又吃瘪
			"abc",
		},

	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, smallestFromLeaf(kit.NewTrees(tc.input).Root), "输入:%v", tc)
	}
}