# [0276-paint-fence](https://leetcode.com/problems/paint-fence)

# 0276.paint fence
## dp framework
1. 设置一个函数,或者 dp 数组,dp(i)表示到第 i 个栅栏,有多少种粉刷方法
2. 递归式 recurrence relation (the hardest one)
3. Base case 0=>0, 1 => k, 2 => k * k

## find recurrence relation
根据题意,::只能有两个连续相同的颜色::,所以我们有两种选择
1. 与上一个栅栏选择不同的颜色,我们知道有 K种颜色,上一个栅栏有dp(i-1)种刷法,所以这里有dp(i-1) * (k-1)刷法
2. 与上一个栅栏选择相同的颜色,所以这里只有1 * dp(i-1)种刷法,但是因为::只能有两个连续相同的颜色::,所以我们选择了和(i-1)栅栏相同颜色之后,必须确保与(i-2)的颜色不同
   那么(i-1)选择与(i-2)不同颜色,有多少种刷法呢?
   => 1 * dp(i-2) *  (k-1)

所以,dp(i) = dp(i-1) * (k-1) + 1 * dp(i-2) *  (k-1)  => (k-1) * (dp(i-1) + dp(i-2))