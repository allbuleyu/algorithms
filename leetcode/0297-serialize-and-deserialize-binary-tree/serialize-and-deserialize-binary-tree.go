package prob0297

import (
	"bytes"
	"algorithms/kit"
	"strconv"
	"strings"
)

type TreeNode = kit.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	b := &bytes.Buffer{}
	doSerializeBfs(root, b)
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

func doSerializeBfs(root *TreeNode, b *bytes.Buffer) {
	queue := make([]*TreeNode, 0)
	var node *TreeNode
	queue = append(queue, root)
	for node != nil || len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		if node == nil {
			b.WriteByte('#')
			b.WriteByte(',')
			continue
		}

		b.WriteString(strconv.Itoa(node.Val))
		b.WriteByte(',')
		queue = append(queue, node.Left, node.Right)
	}
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	d := strings.Split(data, ",")
	return doDeserializeBfs(&d)
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

func doDeserializeBfs(data *[]string) *TreeNode {
	d := *data

	v,_ := strconv.Atoi(d[0])
	root := &TreeNode{Val:v}
	var node *TreeNode
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for i := 1; i < len(d) && len(queue) > 0; {
		node = queue[0]
		queue = queue[1:]

		var left, right *TreeNode
		if d[i] != "#" {
			v1, _ := strconv.Atoi(d[i])
			left = &TreeNode{Val:v1}
			queue = append(queue, left)
		}

		if d[i+1] != "#" {
			v2, _ := strconv.Atoi(d[i+1])
			right = &TreeNode{Val:v2}
			queue = append(queue, right)
		}

		node.Left = left
		node.Right = right

		i += 2
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */