package prob0424

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_characterReplacement(t *testing.T) {
	tcs := []struct{
		input string
		k int

		ans int
	}{
		{
			"ABAB", 2, 4,
		},
		{
			"AABABBA", 1, 4,
		},
	}

	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, characterReplacement(tc.input, tc.k), "输入:%v", tc)
	}
}