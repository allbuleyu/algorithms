package prob0988

import (
	"github.com/allbuleyu/algorithms/kit"
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
func smallestFromLeaf(root *TreeNode) string {
	stringMap := make([]string, 0)
	for i := 'a'; i <= 'z'; i++ {
		stringMap = append(stringMap, string(i))
	}

	return iterate(root, &stringMap)
}

func iterate(root *TreeNode, stringMap *[]string) string {
	if root == nil {
		return ""
	}
	var res string = "z"
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	strStack := make([]string, 0)

	strStack = append(strStack, "")

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ss := strStack[len(strStack)-1]
		strStack = strStack[:len(strStack)-1]

		if node == nil {
			continue
		}

		ss = (*stringMap)[node.Val] + ss
		if node.Left == nil && node.Right == nil {

			if ss < res {
				res = ss
			}
		}

		strStack = append(strStack, ss, ss)
		stack = append(stack, node.Left, node.Right)
	}

	return res
}