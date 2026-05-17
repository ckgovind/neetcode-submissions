func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	ans := 0

	l, r := 0, len(people)-1

	for l <= r {

		if people[r]+people[l] > limit {
			r--
		} else {
			l++
			r--
		}
		ans += 1
	}

	return ans
}
