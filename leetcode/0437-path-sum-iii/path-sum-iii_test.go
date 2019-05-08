package prob0437

import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pathSum(t *testing.T) {

	tcs := []struct{
		input []int
		input1 int
		ans int
	}{
		{
			[]int{10,5,-3,3,2,kit.Null,11,3,-2,kit.Null,1},

			8,
			3,
		},
		{
			[]int{-8,6,8,kit.Null,kit.Null,8,2,kit.Null,kit.Null,kit.Null,-2},

			-2,
			2,
		},
	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, pathSum(kit.NewTrees(tc.input).Root, tc.input1), "输入:%v", tc)
	}
}