package prob0118

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generate(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input int
		ans [][]int
	}{
		//{[]int{10, 5, 2, 6},100,8},
		{5, [][]int{
			[]int{1},
			[]int{1,1},
			[]int{1,2,1},
			[]int{1,3,3,1},
			[]int{1,4,6,4,1},
		}},
		{2, [][]int{
			[]int{1},
			[]int{1,1},
		}},
		{1, [][]int{
			[]int{1},
		}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, generate(tc.input), "输入:%v", tc)
	}
}