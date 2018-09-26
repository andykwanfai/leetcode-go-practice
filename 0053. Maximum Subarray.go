func maxSubArray(nums []int) int {
	//dp solution
	length := len(nums)
	dp := make([]int, length)
	dp[0] = nums[0]
	maximum := dp[0]
	for i := 1; i < length; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maximum = max(maximum, dp[i])
	}
	return maximum
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}