package prob0038

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countAndSay(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input int
		ans string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},

	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, countAndSay(tc.input), "输入:%v", tc)
	}
}

func Test_times(t *testing.T) {
	fmt.Println(countAndSay(12))
}