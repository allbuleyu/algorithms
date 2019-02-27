package prob0746

<<<<<<< HEAD

import "testing"
import (
=======
import (
	"testing"

>>>>>>> 34cfd67db642458eb59534443a215c006677fcca
	"github.com/stretchr/testify/assert"
	"fmt"
)

<<<<<<< HEAD
func Test_prob0236(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		input []int
		ans int
	}{
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},6,},
		{[]int{10, 15, 20},15,},

=======
func Test_prb0001(t *testing.T)  {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans int
	}{
		{
			[]int{10, 15, 20},
			15,
		},
		{
			[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			6,
		},
		{
			[]int{0,0,0,0},
			0,
		},
>>>>>>> 34cfd67db642458eb59534443a215c006677fcca
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, minCostClimbingStairs(tc.input), "输入:%v", tc)
	}
}