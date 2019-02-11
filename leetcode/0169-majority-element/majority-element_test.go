package prob0169

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	ast := assert.New(t)
	tcs := []struct{
		input []int
		ans int
	}{
		{[]int{3,2,3}, 3,},
		{[]int{2,2,1,1,1,2,2}, 2,},
		{[]int{1}, 1},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, majorityElement(tc.input), "输入:%v", tc)
	}
}