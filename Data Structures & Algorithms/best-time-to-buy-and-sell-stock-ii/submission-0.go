func maxProfit(prices []int) int {
	currMinPrice := math.MaxInt

	ans := 0
	for _, val := range prices {
		if currMinPrice < val {
			ans += val - currMinPrice
			// currMinPrice >  val
		}
			currMinPrice = val
	}

	return ans
}
