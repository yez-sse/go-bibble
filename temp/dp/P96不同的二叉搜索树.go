package dp

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	// 这个两重循环的样子，和P343真的是一模一样
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// 1到i为节点组成的二叉搜索树的个数为dp[i]
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
