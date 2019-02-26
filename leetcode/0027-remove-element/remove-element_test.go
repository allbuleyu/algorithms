package prob0027

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_prb0014(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input struct{
			nums []int
			v int
		}
		ans []int
	}{
		{struct {
			nums []int
			v    int
		}{nums: []int{3,2,2,3}, v: 3}, []int{2,2},},
		{struct {
			nums []int
			v    int
		}{nums: []int{}, v: 2}, []int{},},
		{struct {
			nums []int
			v    int
		}{nums: []int{2}, v: 2}, []int{},},
		{struct {
			nums []int
			v    int
		}{nums: []int{1}, v: 2}, []int{1},},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, tc.input.nums[:removeElement(tc.input.nums, tc.input.v)], "输入:%v", tc)
	}
}