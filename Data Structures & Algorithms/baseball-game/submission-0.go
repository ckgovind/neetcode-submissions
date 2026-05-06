
func calPoints(operations []string) int {
	stack := make([]int, 0)
	for _, op := range operations {
		switch op {
		case "+":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = append(stack, a+b)
		case "D":
			a := stack[len(stack)-1]
			stack = append(stack, 2*a)
		case "C":
			stack = stack[:len(stack)-1]
		default:
			val, _ := strconv.Atoi(op)
			stack = append(stack, val)
		}
	}

	s := 0
	for _, val := range stack {
		s += val
	}
	return s
}
