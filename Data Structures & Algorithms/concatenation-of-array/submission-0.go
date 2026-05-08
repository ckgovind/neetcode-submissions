func getConcatenation(nums []int) []int {
    n := len(nums)
	ans := make([]int,2*n)
	for i,_ := range nums{
		ans[i],ans[i+n] = nums[i],nums[i]
	}

	return ans
}
