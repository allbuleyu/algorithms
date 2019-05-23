package prob0107


import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {

	tcs := []struct{
		input []int
		ans [][]int
	}{
		{
			[]int{3,9,20, kit.Null,kit.Null, 15,7},
			[][]int{
				[]int{15,7},
				[]int{9,20},
				[]int{3},
			},
		},
		{
			[]int{1,2,3,4,5,6,7,8},
			[][]int{
				[]int{8},
				[]int{4,5,6,7},
				[]int{2,3},
				[]int{1},
			},
		},

	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, levelOrderBottom(kit.NewTrees(tc.input).Root), "输入:%v", tc)
	}
}