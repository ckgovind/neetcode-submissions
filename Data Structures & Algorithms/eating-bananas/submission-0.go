func maxArr(a []int) int {
	m := 0
	for _, val := range a {
		m = max(m, val)
	}
	return m
}

// check if ans is enough rate to finish within h hours
func isEnough(ans int, h int, piles []int) bool {
	s := 0
	for _, pile := range piles {
		s += (pile + ans - 1) / ans
	}

	fmt.Printf("ISENOUGH FOR %v is %v sum\n", s, h < s)

	return s <=h
}

func minEatingSpeed(piles []int, h int) int {
	end := maxArr(piles) + 1
	start := 1
	ans := 0
	for start < end {
		var mid int = (start + end) / 2
		fmt.Printf("start : %v, end: %v mid: %v\n", start, end, mid)
		if isEnough(mid, h, piles) {
			ans = mid
			end = mid
		} else {
			start = mid + 1
		}
	}

	return ans

}
