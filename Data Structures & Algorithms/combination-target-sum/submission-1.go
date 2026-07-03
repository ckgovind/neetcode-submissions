func sum(arr []int) (s int) {
	for _, val := range arr {
		s += val
	}
	return s
}

func combinationSum(nums []int, target int) [][]int {

	res := make([][]int, 0)
	curr := make([]int, 0)
	var dfs func(int)

	currSum := 0

	dfs = func(indx int) {
		if currSum > target || indx >= len(nums){
			return
		}

		if currSum == target {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		currSum+=nums[indx]
		curr = append(curr, nums[indx])
		dfs(indx)

		currSum-=nums[indx]
		curr = curr[:len(curr)-1]
		dfs(indx + 1)
	}

	dfs((0))

	return res

}
