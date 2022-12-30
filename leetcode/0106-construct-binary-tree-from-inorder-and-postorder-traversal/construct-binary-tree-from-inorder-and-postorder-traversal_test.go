package prob0106

import (
	"fmt"
	"algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildTree(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		inorder []int
		postorder []int
	}{
		{[]int{9,3,15,20,7}, []int{9,15,7,20,3}},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		actual := kit.InOrder(buildTree(tc.inorder, tc.postorder))
		ast.Equal(tc.inorder, actual, "输入:%v", tc)
	}
}