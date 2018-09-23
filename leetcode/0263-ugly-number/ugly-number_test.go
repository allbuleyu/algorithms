package prob0263

import "testing"
import (
	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_prob0236(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		input int
		ans bool
	}{
		{1,true,},
		{2,true,},
		{3,true,},
		{5,true,},
		{7,false,},
		{11,false,},
		{14,false,},
		{31,false,},
		{28,false,},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isUgly(tc.input), "输入:%v", tc)
	}
}