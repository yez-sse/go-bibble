package dp

import "go-learning/util"

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = util.GetMax(curMax, util.GetMax(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}
