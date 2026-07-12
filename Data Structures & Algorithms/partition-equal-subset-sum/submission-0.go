func sum(nums []int) int{
	ans := 0
	for _, val := range nums{
		ans+=val
	}
	return ans
}

func canPartition(nums []int) bool {
    s := sum(nums)
	if s%2 != 0{
		return false
	}

	s = s/2
	dp := make([][]bool,len(nums))

	for i := range dp{
		dp[i] = make([]bool,s+1)
	}


	dp[0][nums[0]] = true

	for i := 1 ; i < len(nums) ; i++{
		for j := 0; j <= s ; j++{
			rem := j - nums[i]
			flag := false
	        if rem >= 0 {
				flag = dp[i-1][rem]
			}

			dp[i][j] = dp[i-1][j] || flag 
		}
	}

	return dp[len(nums)-1][s]
}
