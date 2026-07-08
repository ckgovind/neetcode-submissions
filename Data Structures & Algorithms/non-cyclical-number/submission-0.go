

func isH(n int, seen map[int]bool) bool {

	ans := 0
	k := n
	for k > 0 {
		l := k % 10
		k = k / 10
		ans += (l * l)
	}

	if ans == 1 {
		return true
	} else {
		if seen[ans] {
			return false
		} else {
			seen[ans] = true
			return isH(ans, seen)
		}
	}

}

func isHappy(n int) bool {
	return isH(n, make(map[int]bool))
}
