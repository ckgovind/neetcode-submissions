func maxProfit(prices []int) int {
	stack := make([]int,0)
	ans := 0
	for _, price := range prices{
		for len(stack) > 0 && stack[len(stack)-1] > price{
			ans = max(ans, stack[len(stack)-1] - stack[0])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack,price)
	}
	if len(stack) > 0{
		ans = max(ans, stack[len(stack)-1] - stack[0])
	}
	return ans
}
