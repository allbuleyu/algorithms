package prob0129


import (
	"fmt"
	"algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rightSideView(t *testing.T) {

	tcs := []struct{
		input []int
		ans int
	}{
		{
			[]int{1,2,3},		// 又又吃瘪
			25,
		},
		{
			[]int{4,9,0,5,1},		// 又又吃瘪
			1026,
		},

	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, sumNumbers(kit.NewTrees(tc.input).Root), "输入:%v", tc)
	}
}