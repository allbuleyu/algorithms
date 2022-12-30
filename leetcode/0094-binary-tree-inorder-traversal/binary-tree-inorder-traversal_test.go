package prob0094

import (
	"fmt"
	"algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{

		input, ans []int
	}{
		{
			[]int{1, kit.Null, 2, 3},
			[]int{1,3,2},
		},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, inorderTraversal(kit.NewTrees(tc.input).Head), "输入:%v", tc)
	}
}