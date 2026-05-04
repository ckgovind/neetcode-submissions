func isValid(s string) bool {
	stack := make([]string, 0)
	for _, v := range s {
		val := string(v)
		switch val {
		case "{", "(", "[":
			stack = append(stack, val)
		case "}", ")", "]":
			opp := map[string]string{"}": "{", ")": "(", "]": "["}[val]
			if len(stack) == 0 || stack[len(stack)-1] != opp {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
