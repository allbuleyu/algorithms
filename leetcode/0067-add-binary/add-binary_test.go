package prob0067

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_prob0053(t *testing.T)  {
	ast := assert.New(t)

	tcs := []struct {
		input string
		p2 string
		ans string
	}{
		{"11", "1", "100",},
		{"1010", "1011", "10101",},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, addBinary(tc.input, tc.p2), "输入:%v", tc)
	}
}