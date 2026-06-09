func twoSum(numbers []int, target int) []int {

	startIndx, currIndx := 0, len(numbers)-1

	for startIndx < currIndx {
		currSum := numbers[startIndx] + numbers[currIndx]

		if currSum == target {
			return []int{startIndx+1, currIndx+1}
		} else if currSum > target {
			currIndx--
		} else {
			startIndx++
		}
	}

	return []int{}

}
