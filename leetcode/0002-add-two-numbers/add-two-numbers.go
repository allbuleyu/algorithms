package prob0002

import "github.com/allbuleyu/algorithms/kit"
type ListNode = kit.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return solutionOptimize(l1,l2)
}


// time 100%, space 100%
func solutionOptimize(l1, l2 *ListNode) *ListNode {
	carry := 0
	head := l1

	for l2 != nil {
		l1.Val += l2.Val + carry
		carry = l1.Val / 10
		l1.Val = l1.Val%10

		if l1.Next == nil && carry == 1 {
			l1.Next = &ListNode{Val:0}
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	for carry == 1 {
		l1.Val += carry
		carry = l1.Val/10
		l1.Val %= 10

		if carry == 1 {
			l1.Next = &ListNode{Val:0}
		}

		l1 = l1.Next
	}

	return head
}

// 其实解中sum的空间也可以省略
func solution1(l1, l2 *ListNode) *ListNode {
	carry := 0
	var head = l1
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
		}

		if l2 != nil {
			sum += l2.Val
		}

		sum += carry

		if sum >= 10 {
			carry = 1
		}else {
			carry = 0
		}

		// 把计算结果存入l1中,节省空间
		l1.Val = sum % 10

		if l1.Next == nil {


			// 在l2比l1长的情况下,让l1持续加长,
			//当l2最后一个节点的时候,直接退出,才是正确结果
			if l2 == nil && carry == 0 {
				break
			}

			if l2 != nil && l2.Next == nil && carry == 0 {
				break
			}
			l1.Next = &ListNode{Val:0,Next:nil}
		}



		l1 = l1.Next

		if l2 != nil {
			l2 = l2.Next
		}
	}

	//head = reverseList(head)

	return head
}