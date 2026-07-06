func findMin(nums []int) int {

	first := nums[0]

	start := 0
	last := len(nums) - 1

	ans := first
	for start <= last {
		var mid int = (start + last) / 2
		if nums[mid] >= first {
			start = mid + 1
		} else {
			ans = min(ans,nums[mid])
			last = mid - 1
		}
	}
	return ans

}
