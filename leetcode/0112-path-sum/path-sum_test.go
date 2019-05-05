package prob0112


import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_permute(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct{
		input []int
		input1 int
		ans bool
	}{
		{
		[]int{5,4,8,11,kit.Null,13,4,7,2,kit.Null,kit.Null,kit.Null,kit.Null,kit.Null,1},
		22,
		true,

		},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, hasPathSum(kit.NewTrees(tc.input).Head, tc.input1), "输入:%v", tc)
	}
}