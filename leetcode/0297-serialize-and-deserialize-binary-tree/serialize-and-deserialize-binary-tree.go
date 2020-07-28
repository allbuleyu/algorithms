package prob0297

import (
	"bytes"
	"github.com/allbuleyu/algorithms/kit"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = kit.TreeNode



type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	b := &bytes.Buffer{}
	doSerializePreOrder(root, b)
	ans := b.Bytes()

	return string(ans[:len(ans)-1])
}

func doSerialize(root *TreeNode, do func(root *TreeNode, b *bytes.Buffer)) string {
	b := &bytes.Buffer{}
	do(root, b)

	return b.String()[:b.Len()-1]
}

func doSerializePreOrder(root *TreeNode, b *bytes.Buffer) {
	if root == nil {
		b.WriteString("#")
		b.WriteByte(',')
		return
	}

	b.WriteString(strconv.Itoa(root.Val))
	b.WriteByte(',')

	doSerializePreOrder(root.Left, b)
	doSerializePreOrder(root.Right, b)
}

// can't deserialize! because can't known which is root.
func doSerializeInorder(root *TreeNode, b *bytes.Buffer) {
	if root == nil {
		b.WriteByte('#')
		b.WriteByte(',')
		return
	}

	doSerializeInorder(root.Left, b)
	b.WriteString(strconv.Itoa(root.Val))
	b.WriteByte(',')
	doSerializeInorder(root.Right, b)
}

// result not reverse
func doSerializePostorder(root *TreeNode, b *bytes.Buffer) {
	if root == nil {
		b.WriteByte('#')
		b.WriteByte(',')
		return
	}

	b.WriteString(strconv.Itoa(root.Val))
	b.WriteByte(',')
	doSerializePostorder(root.Right, b)
	doSerializePostorder(root.Left, b)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	d := strings.Split(data, ",")
	return doDeserializePreOrder(&d)
}

func doDeserialize(s string, do func(*[]string) *TreeNode) *TreeNode {
	d := strings.Split(s, ",")
	return do(&d)
}

func doDeserializePreOrder(data *[]string) *TreeNode {
	if len(*data) == 0 {
		return nil
	}

	if (*data)[0] == "#" {
		(*data) = (*data)[1:]
		return nil
	}

	rv, _ := strconv.Atoi((*data)[0])
	root := &TreeNode{Val: rv}
	*data = (*data)[1:]
	root.Left = doDeserializePreOrder(data)
	root.Right = doDeserializePreOrder(data)

	return root
}

func doDeserializePostorder(data *[]string) *TreeNode {
	if len(*data) == 0 {
		return nil
	}

	if (*data)[0] == "#" {
		(*data) = (*data)[1:]
		return nil
	}

	v, _ := strconv.Atoi((*data)[0])
	root := &TreeNode{Val: v}
	(*data) = (*data)[1:]
	root.Right = doDeserializePostorder(data)
	root.Left = doDeserializePostorder(data)

	return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */