
func productExceptSelf(nums []int) []int {
	prefixProd := make([]int, len(nums))
	suffixProd := make([]int, len(nums))

	curr := 1
	for i := 0; i < len(nums); i++ {
		prefixProd[i] = curr * nums[i]
		curr = prefixProd[i]
	}

	curr = 1

	for i := len(nums) - 1; i >= 0; i-- {
		suffixProd[i] = curr * nums[i]
		curr = suffixProd[i]
	}

	ans := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			ans[i] = suffixProd[i+1]
		} else if i == len(nums)-1 {
			ans[i] = prefixProd[i-1]
		} else {
			ans[i] = prefixProd[i-1] * suffixProd[i+1]
		}

	}
	return ans
}
