package 背包

import "go-learning/util"

func canPartition(nums []int) bool {

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	dp := make([]int, 10001)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = util.GetMax(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	if dp[target] == target {
		return true
	}
	return false
}
