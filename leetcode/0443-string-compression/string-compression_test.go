package prob0443

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_compress(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []byte
		ans int
	}{
		{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, 6},
		{[]byte{'a'}, 1},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, compress(tc.input), "输入:%v", tc)
	}
}