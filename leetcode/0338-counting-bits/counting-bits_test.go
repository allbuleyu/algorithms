package prob0338

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countBits(t *testing.T) {
	ast := assert.New(t)

	var tcs = []struct {
		input int
		ans []int
	}{
		{2, []int{0,1,1}},
		{5, []int{0,1,1,2,1,2}},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, countBits(tc.input),"输入:%v", tc)
	}

}