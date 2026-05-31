func abs(a int) int{
	if a < 0 {
		return a * -1
	}
	return a
}

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)
	for i, val := range nums{
		if j, ok := mp[val];ok{
			if abs(i-j) <= k{
				return true
			}
		}
		mp[val] = i
	}

	return false
}
