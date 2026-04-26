
func isAlphaNum(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9')
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b | ' ' // ' ' (space) has the value 32
	}
	return b
}

func isPalindrome(s string) bool {
	sindx := len(s) - 1

	for i := 0; i < sindx; {
		if !isAlphaNum(s[i]) {
			i++
			continue
		}

		if !isAlphaNum(s[sindx]) {
			sindx--
			continue
		}

		if toLower(s[i]) == toLower(s[sindx]) {
			i++
			sindx--
			continue
		} else {
			return false
		}
	}

	return true

}
