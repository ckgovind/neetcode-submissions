func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	first := nums[0]
	if target == first {
		return 0
	}
	for start <= end {
		var mid int = (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= first {
			if target > first {
				// target bigger than first means it will be in the first secion
				// we are in the first section
				if nums[mid] > target {
					end = mid - 1
				} else {
					start = mid + 1
				}
			} else {
				// target is less than first means it will be in the second section
				start = mid + 1
			}
		} else {
			// we are in the second section
			if target > first {
				// we need to move to the first section
				end = mid - 1
			} else {
				// we are in the seond section
				// target is in the second section

				if nums[mid] > target {
					end = mid - 1
				} else {
					start = mid + 1
				}
			}
		}
	}

	return -1
}
