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

	dfs = func(indx int) {
		if sum(curr) > target || indx >= len(nums){
			return
		}

		if sum(curr) == target {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		curr = append(curr, nums[indx])
		dfs(indx)
		curr = curr[:len(curr)-1]
		dfs(indx + 1)
	}

	dfs((0))

	return res

}
