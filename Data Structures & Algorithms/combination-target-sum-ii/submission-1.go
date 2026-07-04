func combinationSum2(nums []int, target int) [][]int {

	res := make([][]int, 0)
	curr := make([]int, 0)
	var dfs func(int)

	currSum := 0

	sort.Ints(nums)
	dfs = func(indx int) {
		// fmt.Printf("The value is %v \n", curr)

		if currSum == target {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		if currSum > target || indx >= len(nums) {
			// fmt.Printf("The leaf value is %v \n", curr)
			return
		}

		currSum += nums[indx]
		curr = append(curr, nums[indx])
		dfs(indx + 1)

		curr = curr[:len(curr)-1]
		currSum -= nums[indx]

		for indx+1 < len(nums) && nums[indx] == nums[indx+1] {
			indx++
		}
		dfs(indx + 1)

	}

	dfs((0))

	return res

}
