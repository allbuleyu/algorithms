package prob0198

func dynamicProgramming(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	for i := 1; i < len(nums);i++ {
		dp[i+1] = max(dp[i], dp[i-1] + nums[i])
	}

	return dp[len(nums)]
}

// 这个老哥的思路也很好,记录!
//First, we consider there are two possible situation: 1. The previous house has been robbed 2. The previous house has not been robbed.
//
//If the previous house has been robbed, we will not be able to rob the current house. So we need to skip this house.
//
//If the previous house has not been robbed, we may rob this current house, or we skip this house.
func rob(nums []int) int {

	var prevNo, prevYes int

	for _, v := range nums {
		var tmp int = prevNo

		prevNo = max(prevNo, prevYes)	// don't rob current house

		prevYes = tmp + v				//rob the current house
	}

	return max(prevNo, prevYes)
}



func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 推导过程
// input x[1...n]
// 设p[1...k] 为抢劫所获最大利益的集合, tips(有i 1 <= i < k => p[i] <p[i+1])
// 那么必有 k = n
// maxPrice = p[k]
// 因为不能抢劫相邻的房子
// 所以:
// 在忽略基本条件的情况下 p[k] = max(p[k-1],p[k-2]+x[n]), 因问题的最优解包含子问题的最优解,以及问题的子问题与子问题的子问题有重叠,用dp
// 基本条件:
// p[0] = x[0]
// 根据tips有: p[1] = max(x[0], x[1])
func robMy(nums []int) int {
	n := len(nums)
	var p []int = make([]int, n)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	p[0] = nums[0]
	p[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		p[i] = max(p[i-1], p[i-2] + nums[i])
	}

	return p[n-1]
}

// 因为求最大利益的过程p[k]只用到了p[k-1], p[k-2]
// 所以可以用两个变量代替数组,节省空间
// https://github.com/aQuaYi/LeetCode-in-Go/tree/master/Algorithms/0198.house-robber
// 根据这位仁兄的答案优化
// k-1, k-2 中,两个数必然有一个是奇数,另外一个肯定为偶数
func robMyOptimize(nums []int) int {
	var odd, even int
	for i, v := range nums {
		if i % 2 == 0 {
			even = max(even+v, odd)
		}else {
			odd = max(even, odd+v)
		}
	}

	return max(odd, even)
}


