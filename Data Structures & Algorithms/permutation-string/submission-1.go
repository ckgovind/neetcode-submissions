
func isSame(a, b []int) bool {

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func checkInclusion(s1 string, s2 string) bool {

	if len(s2) < len(s1) {
		return false
	}

	freqArray := make([]int, 26)
	freqArray2 := make([]int, 26)

	for _, val := range s1 {
		freqArray[val-'a'] += 1
	}

	leftIndx := 0
	rightIndx := 0

	for rightIndx = 0; rightIndx < len(s1); rightIndx++ {
		freqArray2[s2[rightIndx]-'a'] += 1
	}

	if isSame(freqArray, freqArray2) {
		return true
	}

	for rightIndx < len(s2) {
		// remove left index
		freqArray2[s2[rightIndx]-'a'] += 1
		freqArray2[s2[leftIndx]-'a'] -= 1

		if isSame(freqArray, freqArray2) {
			return true
		}

		rightIndx += 1
		leftIndx += 1
	}

	return false
}
