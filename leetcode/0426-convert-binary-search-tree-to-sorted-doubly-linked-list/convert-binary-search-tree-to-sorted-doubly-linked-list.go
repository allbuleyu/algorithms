package prob0426

import "algorithms/kit"

type Node = kit.TreeNode
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return root
	}

	return helper(root)
}

// 用莫里斯中序遍历法
func helpMorris(root *Node) *Node {
	node := root
	head := &Node{Right: nil}
	var p *Node

	for node != nil {
		if node.Left != nil {
			prev := node.Left
			for prev.Right != nil && prev.Right != node {
				prev = prev.Right
			}

			if prev.Right == node {
				if p != nil && p != prev {
					p.Right = prev
					prev.Left = p
				}
				p = node

				node.Left = prev
				node = node.Right
			}else {
				prev.Right = node
				node = node.Left
			}

		}else {
			if head.Right == nil {
				head.Right = node
			}

			if p != nil {
				node.Left = p
				p.Right = node
			}

			p = node
			node = node.Right
		}
	}

	if head.Right != nil {
		head.Right.Left = p
		p.Right = head.Right
	}

	return head.Right
}

// 递归方式处理,还挺有意思的
var head, tail *Node
func helper(root *Node) *Node {
	// 因为是全局变量,所以在执行test case 的时候,要把head, tail重新初始化一下
	head = nil
	tail = nil

	helpRecursion(root)
	head.Left = tail
	tail.Right = head

	return head
}

func helpRecursion(root *Node) {
	if root == nil {
		return
	}

	helpRecursion(root.Left)

	if tail == nil {
		head = root
	}else {
		tail.Right = root
		root.Left = tail
	}

	tail = root
	helpRecursion(root.Right)
}
