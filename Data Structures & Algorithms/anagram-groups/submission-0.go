
func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}

func groupAnagrams(strs []string) [][]string {
	gaMap := make(map[string]int)

	answer := make([][]string, 0)
	i := 0
	for _, str := range strs {
		l := sortString(str)
		if indx, ok := gaMap[l]; ok {
			answer[indx] = append(answer[indx], str)
		} else {
			answer = append(answer, []string{str})
			gaMap[l] = i
			i += 1
		}
	}
	return answer

}
