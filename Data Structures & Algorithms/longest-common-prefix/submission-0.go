
func getCommonPrefix(a string, b []rune) []rune {
	var ans []rune
	for i, val := range b {
		if i >= len(a) {
			return ans
		}
		if val == rune(a[i]) {
			ans = append(ans, val)
		}else{
			return ans
		}
	}
	return ans
}

func longestCommonPrefix(strs []string) string {
	var ans = make([]rune, 0)
	ans = []rune(strs[0])
	for _, str := range strs {
		ans = getCommonPrefix(str, ans)
	}

	return string(ans)
}
