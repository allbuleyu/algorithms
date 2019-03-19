package prob0784

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_letterCasePermutation(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct{
		input string
		ans []string
	}{
		{"a1b2", []string{"a1b2", "a1B2", "A1b2", "A1B2"}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, letterCasePermutation(tc.input), "输入:%v", tc)
	}
}