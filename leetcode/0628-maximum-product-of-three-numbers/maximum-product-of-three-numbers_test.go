package prob0628

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_maximumProduct(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans int
	}{
		{[]int{1,2,3},6},
		{[]int{3,2,1},6},
		{[]int{1,2,3,4},24},
		{[]int{-1,-2,3,4},8},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, maximumProduct(tc.input), "输入:%v", tc)
	}
}