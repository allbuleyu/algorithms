package prob0557

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input string
		ans string
	}{
		{"hello", "olleh"},
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, reverseWords(tc.input), "输入:%v", tc)
	}
}