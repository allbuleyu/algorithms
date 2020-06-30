package prob0100


import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pathSum(t *testing.T) {
	fmt.Println('(', ')')
	return
	tcs := []struct{
		input []int
		input1 []int
		ans bool
	}{
		{
			[]int{1,2,3},

			[]int{1,2,3},
			true,
		},
		{
			[]int{1,1,2},

			[]int{1,2,1},
			false,
		},

		{
			[]int{1,2},

			[]int{1,kit.Null,2},
			false,
		},
		{
			[]int{0,-5},

			[]int{0,8},
			false,
		},
		{
			[]int{1},

			[]int{1},
			true,
		},
	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isSameTree(kit.NewTrees(tc.input).Root, kit.NewTrees(tc.input1).Root), "输入:%v", tc)
	}
}

