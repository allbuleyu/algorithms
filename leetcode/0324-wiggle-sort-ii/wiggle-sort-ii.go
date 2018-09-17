package prob0324

import (
	"sort"
)

//https://leetcode.com/problems/wiggle-sort-ii/description/
//Given an unsorted array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....
//
//Example:
//(1) Given nums = [1, 5, 1, 1, 6, 4], one possible answer is [1, 4, 1, 5, 1, 6].
//(2) Given nums = [1, 3, 2, 2, 3, 1], one possible answer is [2, 3, 1, 3, 1, 2].
//[1,1,2,2,3,3], [1,2,1,3,2,3]
//[1,1,1,4,5,6], [1,4,1,5,1,6]
//例:[1, 5, 1, 1, 6, 4, 7]
//0,1,2,3,4,5,6, 0,4,1,5,2,6,3
//[1,1,1,4,5,6,7], [1,5,1,6,1,7,4]
//Note:
//You may assume all input has valid answer.
//
//Follow Up:
//Can you do it in O(n) time and/or in-place with O(1) extra space?
//
//Credits:
//Special thanks to @dietpepsi for adding this problem and creating all test cases.
//Next challenge 75
func wiggleSort(nums []int)  {
	sorted := sort.IsSorted(sort.IntSlice(nums))
	if !sorted {
		sort.Ints(nums)
	}

	cp := make([]int, len(nums))
	copy(cp, nums)

	left, right := (len(nums) + 1) / 2 - 1, len(nums) - 1
	for i := range nums {
		if i % 2 == 0 {
			nums[i] = cp[left]
			left--
		}else {
			nums[i] = cp[right]
			right--
		}
	}
}

type ss []int
func(n ss)Less(i,j int)bool{
	return n[i]<n[j]
}
func(n ss)Swap(i,j int){
	n[i],n[j] = n[j],n[i]
}
func (n ss)Len()int{
	return len(n)
}
func wiggleSort1(nums []int){
	t := make(ss,len(nums))
	for i:=0;i<len(t);i++{
		t[i] = nums[i]
	}
	sort.Sort(t)
	j:=len(nums) -1
	for i:=1;i<len(nums);i,j=i+2,j-1{
		nums[i] = t[j]
	}
	for i:=0;i<len(nums);i,j=i+2,j-1{
		nums[i] = t[j]
	}

}

func wiggleSort2(nums []int)  {
	if len(nums) < 2 { return }
	newNums := make([]int, len(nums))
	copy(newNums, nums)
	sort.Ints(newNums)
	l, r := (len(nums) + 1) / 2 - 1, len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if i % 2 == 0 {
			nums[i] = newNums[l]
			l--
		} else {
			nums[i] = newNums[r]
			r--
		}
	}
}

func WiggleSort(nums []int) []int {
	wiggleSort(nums)
	return nums
}