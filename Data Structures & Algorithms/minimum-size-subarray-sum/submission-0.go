func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt
	curr := 0

	left, right := 0, 0

	for right = 0; right < len(nums); right++ {
		curr += nums[right]
		// fmt.Printf("The curr is %v\n", curr)
		if curr >= target {
			for left < right {
				// fmt.Printf("The inner loop, the left is %v right is %v and the curr is %v\n", left, right, curr)
				if curr-nums[left] >= target {
					curr = curr - nums[left]
					left++
				} else {
					break
				}
			}
		}
		if curr >= target {
			ans = min(ans, right-left+1)
		}
	}
	if ans == math.MaxInt{
		ans = 0
	}
	return ans
}
