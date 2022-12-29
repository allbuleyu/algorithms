package prob0297

import (
	"fmt"
	"algorithms/kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodec_serialize(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans string
		dAns []int
	}{
		{[]int{1,2,3,kit.Null,kit.Null,4,5},"1,2,#,#,3,4,#,#,5,#,#", []int{1,2,3,4,5}},
	}

	codec := Constructor()
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		tree := kit.NewTrees(tc.input)

		ast.Equal(tc.ans, codec.serialize(tree.Root), "输入:%v", tc)
		root := codec.deserialize(tc.ans)
		ast.Equal(tc.dAns, kit.PreOrder(root), "输入:%v", tc)
	}
}

func TestCodec_serializeInOrder(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans string
		dAns []int
	}{
		{[]int{5,1,4,kit.Null,kit.Null,2,3},"5,4,3,#,#,2,#,#,1,#,#", []int{1,2,3,4,5}},
	}


	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		tree := kit.NewTrees(tc.input)

		ast.Equal(tc.ans, doSerialize(tree.Root, doSerializePostorder), "输入:%v", tc)
		//root := codec.deserialize(tc.ans)
		//ast.Equal(tc.dAns, kit.PostOrder(doDeserialize(tc.ans, doDeserializePostorder)), "输入:%v", tc)
	}
}

func TestCodec_SerializeBfs(t *testing.T) {
	ast := assert.New(t)

	// test case
	tcs := []struct{
		input []int
		ans string
		dAns []int
	}{
		{[]int{1,2,3,kit.Null,kit.Null,4,5},"1,2,3,#,#,4,5,#,#,#,#", []int{1,2,3,4,5}},
	}

	levelOrder := func(root *TreeNode) []int {
		queue := make([]*TreeNode, 0)
		ans := make([]int, 0)
		queue = append(queue, root)
		var node *TreeNode
		for len(queue) > 0 {
			node = queue[0]
			queue = queue[1:]
			ans = append(ans, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		return ans
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		tree := kit.NewTrees(tc.input)

		ast.Equal(tc.ans, doSerialize(tree.Root, doSerializeBfs), "输入:%v", tc)
		//root := codec.deserialize(tc.ans)
		ast.Equal(tc.dAns, levelOrder(doDeserialize(tc.ans, doDeserializeBfs)), "输入:%v", tc)
	}
}