package prob0113


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
		ans [][]int
	}{
		{
			[]int{5,4,8,11,kit.Null,13,4,7,2, kit.Null, kit.Null,5,1},

			22,
			[][]int{
				[]int{5,4,11,2},
				[]int{5,8,4,5},
			},
		},
		{
			[]int{1,2,3,1},
			4,
			[][]int{
				[]int{1,2,1},
				[]int{1,3},
			},
		},

	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, pathSum(kit.NewTrees(tc.input).Root, tc.input1), "输入:%v", tc)
	}
}