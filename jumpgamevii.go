// https://leetcode.com/problems/jump-game-vii/description/
package main

func canReachVII(s string, minJump int, maxJump int) bool {
	l := len(s)
	dp := make([]bool, l)
	count := 0
	dp[0] = true
	for i := 1; i < l; i++ {
		if i >= minJump && dp[i-minJump] {
			count++
		}
		if i >= maxJump+1 && dp[i-maxJump-1] {
			count--
		}
		dp[i] = s[i] == '0' && count > 0

	}
	return dp[l-1]
}
