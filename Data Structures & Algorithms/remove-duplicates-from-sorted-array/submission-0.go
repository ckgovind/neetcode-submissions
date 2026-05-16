
func removeDuplicates(nums []int) int {
	insPointer := 1
	currPointer := 1

	for currPointer < len(nums) {
		if nums[currPointer] != nums[insPointer-1] {
			nums[insPointer] = nums[currPointer]
			currPointer++
			insPointer++
		} else {
			currPointer++
		}
	}
	return insPointer
}
