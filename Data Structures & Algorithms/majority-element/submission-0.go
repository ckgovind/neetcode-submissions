func majorityElement(nums []int) int {
	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num] += 1
		if countMap[num] > len(nums)/2 {
			return num
		}
	}
	return 0
}