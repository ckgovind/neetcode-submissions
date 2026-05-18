func lengthOfLongestSubstring(s string) int {

	indxMap := make(map[rune]int)

	l := 0
	ans := 0
	for i, r := range s {
		if lastIndx, ok := indxMap[r]; ok {
			if lastIndx >= l {
				l = lastIndx + 1
			}
		}
		indxMap[r] = i
		// fmt.Printf("Current values are l:%v i:%v, i-l+1 : %v\n", l, i, i-l+1)
		ans = max(ans, i-l+1)
	}

	ans = max(ans, len(s)-l)

	return ans
}
