package prob0345

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseVowels(t *testing.T) {

	tcs := []struct{
		input string

		ans string
	}{
		{"hello", "holle",},
		{"leetcode", "leotcede",},

	}

	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, reverseVowels(tc.input), "输入:%v", tc)
	}
}