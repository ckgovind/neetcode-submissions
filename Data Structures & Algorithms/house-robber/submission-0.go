func rob(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}

	dp := make([]int,len(nums))
	// dp[i] stores the max value that is obtained by robbing houses from 0..i

	dp[0] = nums[0]
	dp[1] = max(nums[0],nums[1])
	for i,val := range nums{
		if i <=1{
			continue
		}

		dp[i] = max(dp[i-1],val + dp[i-2])
	}

	return dp[len(nums)-1]
}
