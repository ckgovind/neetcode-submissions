func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	
	hsh := make(map[rune]int)
	for _, val := range s {
		hsh[val] += 1
	}

	for _, val := range t {
		hsh[val] -= 1
	}

	for _, val := range hsh {
		if val != 0 {
			return false
		}
	}
	return true
}
