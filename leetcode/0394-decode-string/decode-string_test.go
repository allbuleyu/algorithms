package prob0394

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decodeString(t *testing.T) {
	tcs := []struct{
		input string
		ans string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		//{"3[a2[c]]", "accaccacc"},
		//{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}

	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, decodeString(tc.input), "输入:%v", tc)
	}
}