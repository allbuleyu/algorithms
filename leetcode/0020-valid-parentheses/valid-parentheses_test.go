package prob0020


import (
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_prb0014(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input string
		ans bool
	}{
		{"()", true,},
		{"()[]{}", true,},
		{"(}", false,},
		{"()}", false,},
		{"", true,},
		{"[(s)[s]{a}]", true,},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isValid(tc.input), "输入:%v", tc)
	}
}