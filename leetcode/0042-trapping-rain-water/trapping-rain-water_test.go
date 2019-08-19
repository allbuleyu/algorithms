package prob0042

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_trap(t *testing.T) {

	tcs := []struct{
		input []int
		ans int
	}{
		{
			[]int{0,1,0,2,1,0,1,3,2,1,2,1},
			6,
		},

	}


	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, trap(tc.input), "输入:%v", tc)
	}
}