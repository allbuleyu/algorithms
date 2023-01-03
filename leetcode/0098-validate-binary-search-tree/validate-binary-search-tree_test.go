package prob0098

import (
	"algorithms/kit"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_morris(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct {
		input []int
		ans   bool
	}{
		{[]int{-2147483648}, true},
		{[]int{2, 2, 2}, false},
		{[]int{2, 1, 3}, true},
		{[]int{5, 1, 4, kit.Null, kit.Null, 3, 6}, false},
		{[]int{5, 1, 6, kit.Null, kit.Null, 3, 7}, false},
		{[]int{5, 3, 5, 2, 4, 6, 12}, false},
		{[]int{5, 1, 4, kit.Null, kit.Null, 3, 6}, false},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		tree := kit.NewTrees(tc.input)
		ast.Equal(tc.ans, morris(tree.Root), "输入:%v", tc)
	}
}
