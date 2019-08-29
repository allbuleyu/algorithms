package prob0075

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortColors(t *testing.T) {

	tcs := []struct{
		input []int
		ans []int
	}{
		{
			[]int{2,0,2,1,1,0}, []int{0,0,1,1,2,2},
		},
		{
			[]int{2,1,2,0,0,1}, []int{0,0,1,1,2,2},
		},
	}

	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		sortColors(tc.input)
		ast.Equal(tc.ans, tc.input, "输入:%v", tc)
	}
}