package prob0200

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numIslands(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct{
		input [][]byte
		ans int
	}{
		{
			[][]byte{
				[]byte{'1','1','1','1','0'},
				[]byte{'1','1','0','1','0'},
				[]byte{'1','1','0','0','0'},
				[]byte{'0','0','0','0','0'},
			},1,
		},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, numIslands(tc.input), "输入:%v", tc)
	}
}