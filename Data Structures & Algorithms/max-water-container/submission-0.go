func maxArea(heights []int) int {

	i := 0
	j := len(heights) - 1

	var ans int
	for i <= j {
		currArea := min(heights[i], heights[j]) * (j - (i))
		ans = max(currArea, ans)

		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}

	return ans
}
