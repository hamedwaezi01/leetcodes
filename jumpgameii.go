// https://leetcode.com/problems/jump-game-ii/
package main

func jump(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = 10001
	}
	dp[0] = 0 // just to be explicit

	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[j]+j >= i && dp[j]+1 < dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[l-1]
}
