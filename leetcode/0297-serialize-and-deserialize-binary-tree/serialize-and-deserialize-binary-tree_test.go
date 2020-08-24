package prob0297

import (
	"fmt"
	"github.com/allbuleyu/algorithms/kit"
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
		ast.Equal(tc.dAns, kit.PostOrder(doDeserialize(tc.ans, doDeserializePostorder)), "输入:%v", tc)
	}
}