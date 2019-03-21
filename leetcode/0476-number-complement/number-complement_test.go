package prob0476

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_findComplement(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input, ans int
	}{
		{5, 2},
		{1, 0},
		{9, 6},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, findComplement(tc.input), "输入:%v", tc)
	}
}