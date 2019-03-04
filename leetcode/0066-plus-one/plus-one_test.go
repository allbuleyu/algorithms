package prob0066


import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_prob0053(t *testing.T)  {
	ast := assert.New(t)

	tcs := []struct {
		input []int
		ans []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 4},
		},

		{
			[]int{1,9,9},
			[]int{2,0,0},
		},

		{
			[]int{9,9},
			[]int{1,0,0},
		},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, plusOne(tc.input), "输入:%v", tc)
	}
}