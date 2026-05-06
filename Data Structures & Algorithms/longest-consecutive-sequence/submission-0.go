
func longestConsecutive(nums []int) int {

	ans := 0
	lMap := make(map[int]int)
	rMap := make(map[int]int)
	seen := make(map[int]bool)

	for _, num := range nums {
		if s := seen[num]; !s {
			// num can extend a seq in 3 ways
			// num can be the rightmost new addition to a seq
			a := rMap[num-1]
			rMap[num] = max(rMap[num], a+1)
			// num can be the leftmost new addition to a seq
			b := lMap[num+1]
			lMap[num] = max(lMap[num], b+1)
			// num can join 2 big seq into bigger seq
			c := a + b + 1
			lMap[num-a] = max(lMap[num-a], c)
			rMap[num+b] = max(rMap[num+b], c)

			ans = max(ans, a, b, c)

		}
		seen[num] = true
	}
	return ans
}
