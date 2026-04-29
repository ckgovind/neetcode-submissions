
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	quad := []int{}
	res := make([][]int, 0)

	var ksum func(int, int, []int, int)
	ksum = func(k int, start int, nums []int, target int) {
		if k == 2 {
			// 2-sum
			left := start
			right := len(nums) - 1
			for left < right {
				if nums[left]+nums[right] > target {
					right--
				} else if nums[left]+nums[right] < target {
					left++
				} else if nums[left]+nums[right] == target {
					temp := make([]int, len(quad))
					copy(temp, quad)
					temp = append(temp, nums[left], nums[right])
					res = append(res, temp)
					left++
					right--

					for left < right && nums[left] == nums[left-1] {
						left++
					}

					for left < right && nums[right] == nums[right+1] {
						right--
					}
				}
			}

			return
		}

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			quad = append(quad, nums[i])
			ksum(k-1, i+1, nums, target-nums[i])
			quad = quad[:len(quad)-1]
		}
	}

	ksum(4, 0, nums, target)

	return res

}
