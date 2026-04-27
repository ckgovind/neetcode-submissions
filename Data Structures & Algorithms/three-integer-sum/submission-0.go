
func threeSum(nums []int) [][]int {
	numMap := make(map[int]int)

	sort.Ints(nums)

	for _, val := range nums {
		numMap[val] += 1
	}

	res := make([][]int, 0)

	for i := 0; i < len(nums); i++ {

		numMap[nums[i]] -= 1

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			numMap[nums[j]] -= 1

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			val := 0 - (nums[j] + nums[i])
			if k := numMap[val]; k > 0 {
				res = append(res, []int{nums[i], nums[j], val})
			}
		}

		for j := i + 1; j < len(nums); j++ {
			numMap[nums[j]] += 1
		}
	}

	return res
}
