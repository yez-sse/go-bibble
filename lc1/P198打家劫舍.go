package lc1

import "go-learning/util"

func rob(nums []int) int {
	len := len(nums)
	if len == 1 {
		return nums[0]
	}
	dp := make([]int, len)
	dp[0] = nums[0]
	dp[1] = util.GetMax(nums[0], nums[1])
	for i := 2; i < len; i++ {
		dp[i] = util.GetMax(dp[i - 1], dp[i -2] + nums[i])
	}
	return dp[len - 1]
}
