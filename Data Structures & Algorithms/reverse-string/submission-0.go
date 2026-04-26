func reverseString(s []byte) {
	indx := len(s) - 1
	for i := 0; i <= indx; i++ {
		s[i], s[indx] = s[indx], s[i]
		indx--
	}
}
