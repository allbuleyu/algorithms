package prob0142

import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input []int
		pos int
		ans int
	}{
		// TODO: Add test cases.
		{[]int{3,2,0,-4}, 1, 2},
		{[]int{1, 2}, 0, 1},
		{[]int{1}, 0, 1},


	}

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)
		circleNode := kit.CreateNodesWithPosCircle(tc.input, tc.pos)
		ans := detectCycle(circleNode)
		ast.Equal(tc.ans, ans.Val, "输入:%v", tc)
	}
}