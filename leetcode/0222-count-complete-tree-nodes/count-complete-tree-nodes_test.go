package prob0222

import (
	"fmt"
	"algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_smallestFromLeaf(t *testing.T) {

	tcs := []struct{
		input []int
		ans int
	}{
		{
			[]int{1,2,3,4,5},
			6,
		},


	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, countNodes(kit.NewTrees(tc.input).Root), "输入:%v", tc)
	}
}