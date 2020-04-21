package challenge

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	
	return iteration(preorder)
}

func iteration(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:nums[0],
	}

	tmp := make([]*TreeNode, 1)
	tmp[0] = root

	for i := 1; i < len(nums); i++ {
		t := tmp[len(tmp)-1]
		if nums[i] < t.Val {
			t.Left = &TreeNode{
				Val:nums[i],
			}

			tmp = append(tmp, t.Left)
			continue
		}

		for k := 0; k < len(tmp); k++ {
			if tmp[k].Val < nums[i] {
				t := &TreeNode{
					Val:nums[i],
				}
				tmp[k].Right = t

				tmp = tmp[:k]

				tmp = append(tmp, t)
				break
			}
		}
	}

	return root
}

func recursion(root *TreeNode, nums []int, k int) *TreeNode {
	return &TreeNode{}
}