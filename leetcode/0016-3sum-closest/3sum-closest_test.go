package prob0016


import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {

	tcs := []struct{
		input []int
		target int
		ans int
	}{
		{
			[]int{-1, 2, 1, -4},1,2,
		},
		{
			[]int{-1, 2, 1, -4},0,-1,
		},
	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, threeSumClosest(tc.input, tc.target), "输入:%v", tc)
	}
}