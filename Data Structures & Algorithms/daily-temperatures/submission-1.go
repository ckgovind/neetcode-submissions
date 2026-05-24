func dailyTemperatures(temperatures []int) []int {
	st := make([]int, 0)
	ans := make([]int, 0)

	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(st) > 0 && temperatures[st[len(st)-1]] <= temperatures[i] {
			// pop and discard
			st = st[:len(st)-1]
		}

		if len(st) == 0 {
			ans = append([]int{0}, ans...)
		} else {
			ans = append([]int{st[len(st)-1]-i},ans...)
		}

		st = append(st,i)
	}

	return ans
}
