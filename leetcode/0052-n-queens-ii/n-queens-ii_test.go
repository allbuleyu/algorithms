package prob0052

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_totalNQueens(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input int

		ans int
	}{
		{4,2},


	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, totalNQueens(tc.input), "输入:%v", tc)
	}
}