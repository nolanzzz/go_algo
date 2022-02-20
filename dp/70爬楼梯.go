package dp

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	// dp[i]到达第i节有几种方法
	dp := make([]int, n+1)
	// dp[0]不用考虑，不存在这个情况
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
