func subarraySum(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		curr := 0
		for j := i; j < len(nums); j++ {
			// find the sum of the subarray [i..j]
			curr = curr + nums[j]
			if curr == k {
				ans += 1
			}
		}
	}

	return ans
}
