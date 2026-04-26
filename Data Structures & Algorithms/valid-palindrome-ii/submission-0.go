
func isPal(s string) bool {
	if len(s) < 2 {
		return true
	}

	return (s[0] == s[len(s)-1]) && isPal(s[1:len(s)-1])
}

func validPalindrome(s string) bool {
	indx := len(s) - 1
	for i := 0; i <= indx; {
		if s[i] == s[indx] {
			i++
			indx--
		} else {
			// found the mismatch
			// check if ans is s[i] or s[indx]
			return isPal(s[i+1:indx+1]) || isPal(s[i:indx])

		}
	}
	return true
}
