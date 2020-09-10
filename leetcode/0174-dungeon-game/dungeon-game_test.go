package prob0174

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calculateMinimumHP(t *testing.T) {
	tcs := []struct{
		input [][]int
		ans int
	}{
		{[][]int{[]int{-2,-3,3},[]int{-5,-10,1},[]int{10,30,-5}}, 7},
		//{[][]int{[]int{-2,-3,-3},[]int{-5,-10,1},[]int{10,30,-5}}, 8},

	}

	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, calculateMinimumHP(tc.input), "输入:%v", tc)
	}
}