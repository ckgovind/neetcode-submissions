
func majorityElement(nums []int) []int {

	// there can be atmost 2 elements with count >= n/3
	var a, aCount int
	var b, bCount int

	for _, val := range nums {
		if val == a {
			aCount++
		} else if val == b {
			bCount++
		} else {
			if aCount == 0 {
				a = val
				aCount++
			} else if bCount == 0 {
				b = val
				bCount++
			} else {
				aCount--
				bCount--
			}
		}
	}

	thres := len(nums) / 3
	aC := 0
	bC := 0
	for _, val := range nums {
		if val == a {
			aC++
		} else if val == b {
			bC++
		}
	}
	ans := make([]int, 0)
	if aC > thres {
		ans = append(ans, a)
	}
	if bC > thres {
		ans = append(ans, b)
	}

	return ans

}
