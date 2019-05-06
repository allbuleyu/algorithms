package kit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

)

func TestNewTrees(t *testing.T) {
	//tc := []int{1, 2, 3,4,5,6,7, 8, Null,9,9,Null,Null,10}
	//tc := []int{1,Null,2,Null,3}
	tc := []int{1, 2, 3,4,5, Null,6}

	nodes := NewTrees(tc)

	//fmt.Println(nodes.InOrder(), nodes.depth)

	fmt.Println(nodes.String(), nodes.PostOrder())
}

func Test_toIntsPreOrderIterate(t *testing.T) {
	tc := []int{1, 2, 3,4,5, Null,6}
	ans := []int{1,2,4,5,3,6}

	trees := NewTrees(tc)

	ast := assert.New(t)

	ast.Equal(ans, toIntsPreOrderIterate(trees.Root))
}

func Test_toIntsInOrderIterate(t *testing.T) {
	tc := []int{1, 2, 3,4,5, Null,6}
	ans := []int{4,2,5,1,3,6}

	trees := NewTrees(tc)

	ast := assert.New(t)

	ast.Equal(ans, toIntsInOrderIterate(trees.Root))
}


func Test_toIntsPostOrderIterate(t *testing.T) {
	tc := []int{1, 2, 3,4,5, Null,6}
	ans := []int{4,5,2,6,3,1}

	trees := NewTrees(tc)

	ast := assert.New(t)

	ast.Equal(ans, toIntsPostOrderIterate(trees.Root))
}