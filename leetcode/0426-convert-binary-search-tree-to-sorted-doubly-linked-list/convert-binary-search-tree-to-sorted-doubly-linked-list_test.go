package prob0426

import (
	"fmt"
	"algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_treeToDoublyList(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct{
		input []int
		ans []int
	}{
		{[]int{4,2,5,1,3},[]int{1,2,3,4,5}},
		//{[]int{2,1},[]int{1,2}},
		{[]int{2,1,3},[]int{1,2,3}},
		//{[]int{},[]int{}},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		x := kit.NewTrees(tc.input)
		fmt.Println(x.String())
		a := treeToDoublyList(x.Root)
		head := a

		ans := make([]int, 0)
		for head != nil {
			ans = append(ans, head.Val)
			head = head.Right
			if head == a {
				break
			}
		}
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}