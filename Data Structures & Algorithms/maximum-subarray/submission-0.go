func maxSubArray(nums []int) int {

	currSum := 0
	maxSum := nums[0]

	for _, val := range nums {

		currSum = max(currSum, 0)

		currSum += val
		maxSum = max(currSum, maxSum)
	}

	return maxSum
}
