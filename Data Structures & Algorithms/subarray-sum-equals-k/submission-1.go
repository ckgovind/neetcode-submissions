func subarraySum(nums []int, k int) int {
	ans := 0
	curr := 0
	prefixSum := map[int]int{0: 1}
	for _, val := range nums {
		curr += val
		ans += prefixSum[curr-k]
		prefixSum[curr] += 1
	}
	return ans
}
