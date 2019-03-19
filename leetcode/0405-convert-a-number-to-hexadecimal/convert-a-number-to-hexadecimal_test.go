package prob0405

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_toOctonary(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct{
		input int
		ans string
	}{
		{8, "10"},
		{9, "11"},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, toOctonary(tc.input), "输入:%v", tc)
	}
}