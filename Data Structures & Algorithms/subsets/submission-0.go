func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	curr := nums[0]

	newnums := nums[1:]

	allsubs := subsets(newnums)

	ans := make([][]int, 0)
	for _, subset := range allsubs {
		n := append([]int{}, subset...)
		n = append(n, curr)
		ans = append(ans, n)
		ans = append(ans, subset)
	}

	return ans

}
