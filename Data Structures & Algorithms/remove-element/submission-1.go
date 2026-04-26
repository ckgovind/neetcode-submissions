
func removeElement(nums []int, val int) int {

	indx := len(nums) - 1

	for i := 0; i <= indx; {
		if nums[i] == val {
			nums[i], nums[indx] = nums[indx], nums[i]
			indx--
		}else{
		i++
		}
	}

	return indx + 1
}
